// utf.go.

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

// Unicode Symbols Processing.

package unicode

import (
	"math"
	"unicode/utf8"
)

// CreateValidUtf8Runes Function creates a full Set of available Unicode UTF-8
// Symbols.
func CreateValidUtf8Runes() []rune {

	var r rune
	var runes = make([]rune, 0)

	r = 0
	for {
		if utf8.ValidRune(r) {
			runes = append(runes, r)
		}

		// Next Rune.
		if r == math.MaxInt32 {
			break
		}
		r++
	}

	return runes
}
