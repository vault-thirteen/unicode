// symbol.go.

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

// SymbolIsRusLatLetter Function checks whether the specified Symbol is a Letter
// of Russian or Latin Alphabets or not.
func SymbolIsRusLatLetter(symbol rune) bool {
	if SymbolIsLatLetter(symbol) {
		return true
	}
	if SymbolIsRusLetter(symbol) {
		return true
	}
	return false
}

// SymbolIsLatLetter Function checks whether the specified Symbol is a Letter
// of Latin Alphabet or not.
func SymbolIsLatLetter(symbol rune) bool {
	if (symbol >= 'a') && (symbol <= 'z') {
		return true
	}
	if (symbol >= 'A') && (symbol <= 'Z') {
		return true
	}
	return false
}

// SymbolIsRusLetter Function checks whether the specified Symbol is a Letter
// of Russian Alphabet or not.
func SymbolIsRusLetter(symbol rune) bool {
	if (symbol >= 'а') && (symbol <= 'я') {
		return true
	}
	if (symbol >= 'А') && (symbol <= 'Я') {
		return true
	}
	if symbol == 'ё' {
		return true
	}
	if symbol == 'Ё' {
		return true
	}
	return false
}

// SymbolIsNumber Function checks whether the specified Symbol is numeric.
func SymbolIsNumber(symbol rune) bool {
	if (symbol >= '0') && (symbol <= '9') {
		return true
	}
	return false
}
