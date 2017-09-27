/*
 * Copyright (c) 1987, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)sysexits.h	8.1 (Berkeley) 6/2/93
 */

// Package sysexits is a go port of sysexits.h from glibc. See `man sysexits`.
package sysexits

const (
	// OK (0) is the code for successful termination.
	OK = 0

	// Usage (64) indicates that the command was used incorrectly, e.g., with
	// the wrong number of arguments, a bad flag, a bad syntax in a parameter,
	// or whatever.
	Usage = 64

	// DataErr (65) indicates that the input data was incorrect in some way.
	// This should only be used for user's data and not system files.
	DataErr = 65

	// NoInput (66) indicates that an input file (not a system file) did not
	// exist or was not readable.  This could also include errors like ``No
	// message'' to a mailer (if it cared to catch it).
	NoInput = 66

	// NoUser (67) indicates that the user specified did not exist. This might
	// be used for mail addresses or remote logins.
	NoUser = 67

	// NoHost (68) indicates that the host specified did not exist.  This is
	// used in mail addresses or network requests.
	NoHost = 68

	// Unavailable (69) indicates that a service is unavailable.  This can
	// occur if a support program or file does not exist. This can also be used
	// as a catchall message when something you wanted to do doesn't work, but
	// you don't know why.
	Unavailable = 69

	// Software (70) indicates that an internal software error has been
	// detected. This should be limited to non-operating system related errors
	// as possible.
	Software = 70

	// OSErr (71) indicates that an operating system error has been detected.
	// This is intended to be used for such things as ``cannot fork'', ``cannot
	// create pipe'', or the like.  It includes things like getuid returning a
	// user that does not exist in the passwd file.
	OSErr = 71

	// OSFile (72) indicates that some system file (e.g., /etc/passwd,
	// /var/run/utmp, etc.) does not exist, cannot be opened, or has some sort
	// of error (e.g., syntax error).
	OSFile = 72

	// CantCreat (73) indicates that a (user specified) output file cannot be
	// created.  Why it's called creat exactly, I don't know. I suppose it's
	// faster to say. See `man 2 creat`.
	CantCreat = 73

	// IOErr (74) indicates that an error occurred while doing I/O on some
	// file.
	IOErr = 74

	// TempFail (75) is the code for temporary failure, indicating something
	// that is not really an error. In sendmail, this means that a mailer
	// (e.g.) could not create a connection, and the request should be
	// reattempted later.
	TempFail = 75

	// Protocol (76) indicates that the remote system returned something that
	// was ``not possible'' during a protocol exchange.
	Protocol = 76

	// NoPerm (77) indicates that you did not have sufficient permission to
	// perform the operation.  This is not intended for file system problems,
	// which should use NoInput or CantCreat, but rather for higher level
	// permissions.
	NoPerm = 77

	// Config (78) indicates that something was found in an unconfigured or
	// misconfigured state.
	Config = 78
)
