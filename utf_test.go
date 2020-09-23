// utf_test.go.

// +build test

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package unicode

import (
	"os"
	"testing"
	"time"
	"unicode/utf8"
)

func Test_CreateUtf8Runes(t *testing.T) {

	const tmpFileName = "/tmp/runes.tmp.txt"

	var err error
	var file *os.File
	var runes []rune

	file, err = os.OpenFile(tmpFileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.FailNow()
	}

	runes = CreateValidUtf8Runes()
	for _, aRune := range runes {
		if !utf8.ValidRune(aRune) {
			t.FailNow()
		}
		_, err = file.Write([]byte(string(aRune)))
		if err != nil {
			t.FailNow()
		}
	}
	if len(runes) != 1112064 {
		t.FailNow()
	}

	err = file.Close()
	if err != nil {
		t.FailNow()
	}

	time.Sleep(time.Second * 1) // For manual Check.

	err = os.Remove(tmpFileName)
	if err != nil {
		t.FailNow()
	}
}
