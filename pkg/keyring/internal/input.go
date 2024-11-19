package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bgentry/speakeasy"
	"github.com/mattn/go-isatty"
	"golang.org/x/crypto/bcrypt"
)

const (
	maxPassphraseEntryAttempts = 3
	// MinPassLength is the minimum acceptable password length
	MinPassLength = 8
)

// NewRealPrompt creates a function that prompts for and manages keyring passphrases.
func NewRealPrompt(dir string, buf io.Reader) func(string) (string, error) {
	return func(prompt string) (string, error) {
		keyhashStored := false
		keyhashFilePath := filepath.Join(dir, "keyhash")

		var keyhash []byte

		_, err := os.Stat(keyhashFilePath)

		switch {
		case err == nil:
			keyhash, err = os.ReadFile(keyhashFilePath)
			if err != nil {
				return "", fmt.Errorf("failed to read %s: %w", keyhashFilePath, err)
			}

			keyhashStored = true

		case os.IsNotExist(err):
			keyhashStored = false

		default:
			return "", fmt.Errorf("failed to open %s: %w", keyhashFilePath, err)
		}

		failureCounter := 0

		for {
			failureCounter++
			if failureCounter > maxPassphraseEntryAttempts {
				return "", fmt.Errorf("too many failed passphrase attempts")
			}

			buf := bufio.NewReader(buf)
			pass, err := getPassword(fmt.Sprintf("Enter keyring passphrase (attempt %d/%d):", failureCounter, maxPassphraseEntryAttempts), buf)
			if err != nil {
				// NOTE: LGTM.io reports a false positive alert that states we are printing the password,
				// but we only log the error.
				//
				// lgtm [go/clear-text-logging]
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			if keyhashStored {
				if err := bcrypt.CompareHashAndPassword(keyhash, []byte(pass)); err != nil {
					fmt.Fprintln(os.Stderr, "incorrect passphrase")
					continue
				}

				return pass, nil
			}

			reEnteredPass, err := getPassword("Re-enter keyring passphrase:", buf)
			if err != nil {
				// NOTE: LGTM.io reports a false positive alert that states we are printing the password,
				// but we only log the error.
				//
				// lgtm [go/clear-text-logging]
				_, _ = fmt.Fprintln(os.Stderr, err)
				continue
			}

			if pass != reEnteredPass {
				_, _ = fmt.Fprintln(os.Stderr, "passphrase do not match")
				continue
			}

			passwordHash, err := bcrypt.GenerateFromPassword([]byte(pass), 2)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				continue
			}

			if err := os.WriteFile(keyhashFilePath, passwordHash, 0o600); err != nil {
				return "", err
			}

			return pass, nil
		}
	}
}

// getPassword will prompt for a password one-time
// It enforces the password length
func getPassword(prompt string, buf *bufio.Reader) (pass string, err error) {
	if inputIsTty() {
		pass, err = speakeasy.FAsk(os.Stderr, prompt)
	} else {
		pass, err = readLineFromBuf(buf)
	}

	if err != nil {
		return "", err
	}

	if len(pass) < MinPassLength {
		// Return the given password to the upstream client so it can handle a
		// non-STDIN failure gracefully.
		return pass, fmt.Errorf("password must be at least %d characters", MinPassLength)
	}

	return pass, nil
}

// inputIsTty returns true iff we have an interactive prompt,
// where we can disable echo and request to repeat the password.
// If false, we can optimize for piped input from another command
func inputIsTty() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
}

// readLineFromBuf reads one line from reader.
// Subsequent calls reuse the same buffer, so we don't lose
// any input when reading a password twice (to verify)
func readLineFromBuf(buf *bufio.Reader) (string, error) {
	pass, err := buf.ReadString('\n')

	switch {
	case errors.Is(err, io.EOF):
		// If by any chance the error is EOF, but we were actually able to read
		// something from the reader then don't return the EOF error.
		// If we didn't read anything from the reader and got the EOF error, then
		// it's safe to return EOF back to the caller.
		if len(pass) > 0 {
			// exit the switch statement
			break
		}
		return "", err

	case err != nil:
		return "", err
	}

	return strings.TrimSpace(pass), nil
}
