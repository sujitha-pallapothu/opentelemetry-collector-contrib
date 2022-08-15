// Code generated from Sql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var sqllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqllexerLexerInit() {
	staticData := &sqllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "','", "'('", "')'", "';'", "", "", "", "", "", "", "",
		"", "", "", "'='", "'>'", "'<'", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "'*'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "WS", "COMMA", "L_BRACKET", "R_BRACKET", "EOQ", "BOOLEAN_LITERAL",
		"K_SELECT", "K_WHERE", "K_WINDOW_TUMBLING", "K_GROUP_BY", "K_AND", "K_OR",
		"K_IS", "K_LIKE", "K_NOT_LIKE", "K_EQUAL", "K_GREATER", "K_LESS", "K_LESS_EQUAL",
		"K_GREATER_EQUAL", "K_NOT_EQUAL", "K_NULL", "K_IS_NULL", "K_IS_NOT_NULL",
		"K_NOT", "K_NOT_IN", "K_IN", "K_COUNT", "K_SUM", "K_MIN", "K_MAX", "K_AVG",
		"K_TRUE", "K_FALSE", "K_UPPER", "K_LOWER", "IDENTIFIER", "NUMERIC_LITERAL",
		"STRING_LITERAL", "STAR",
	}
	staticData.ruleNames = []string{
		"SPACE", "WS", "COMMA", "L_BRACKET", "R_BRACKET", "EOQ", "BOOLEAN_LITERAL",
		"K_SELECT", "K_WHERE", "K_WINDOW_TUMBLING", "K_GROUP_BY", "K_AND", "K_OR",
		"K_IS", "K_LIKE", "K_NOT_LIKE", "K_EQUAL", "K_GREATER", "K_LESS", "K_LESS_EQUAL",
		"K_GREATER_EQUAL", "K_NOT_EQUAL", "K_NULL", "K_IS_NULL", "K_IS_NOT_NULL",
		"K_NOT", "K_NOT_IN", "K_IN", "K_COUNT", "K_SUM", "K_MIN", "K_MAX", "K_AVG",
		"K_TRUE", "K_FALSE", "K_UPPER", "K_LOWER", "IDENTIFIER", "NUMERIC_LITERAL",
		"STRING_LITERAL", "STAR", "DIGIT", "A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 457, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 143, 8, 1, 11, 1, 12, 1, 144, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6,
		159, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 314, 8, 37, 10, 37, 12, 37, 317, 9,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 324, 8, 37, 10, 37, 12, 37,
		327, 9, 37, 1, 37, 1, 37, 1, 37, 5, 37, 332, 8, 37, 10, 37, 12, 37, 335,
		9, 37, 1, 37, 1, 37, 1, 37, 5, 37, 340, 8, 37, 10, 37, 12, 37, 343, 9,
		37, 3, 37, 345, 8, 37, 1, 38, 4, 38, 348, 8, 38, 11, 38, 12, 38, 349, 1,
		38, 1, 38, 5, 38, 354, 8, 38, 10, 38, 12, 38, 357, 9, 38, 3, 38, 359, 8,
		38, 1, 38, 1, 38, 3, 38, 363, 8, 38, 1, 38, 4, 38, 366, 8, 38, 11, 38,
		12, 38, 367, 3, 38, 370, 8, 38, 1, 38, 1, 38, 4, 38, 374, 8, 38, 11, 38,
		12, 38, 375, 1, 38, 1, 38, 3, 38, 380, 8, 38, 1, 38, 4, 38, 383, 8, 38,
		11, 38, 12, 38, 384, 3, 38, 387, 8, 38, 3, 38, 389, 8, 38, 1, 39, 1, 39,
		1, 39, 1, 39, 5, 39, 395, 8, 39, 10, 39, 12, 39, 398, 9, 39, 1, 39, 1,
		39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1,
		50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55,
		1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1,
		60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65,
		1, 66, 1, 66, 1, 67, 1, 67, 0, 0, 68, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 0,
		85, 0, 87, 0, 89, 0, 91, 0, 93, 0, 95, 0, 97, 0, 99, 0, 101, 0, 103, 0,
		105, 0, 107, 0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119, 0, 121, 0,
		123, 0, 125, 0, 127, 0, 129, 0, 131, 0, 133, 0, 135, 0, 1, 0, 36, 3, 0,
		9, 11, 13, 13, 32, 32, 2, 0, 9, 9, 32, 32, 1, 0, 34, 34, 1, 0, 96, 96,
		1, 0, 93, 93, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 2, 0, 43, 43, 45, 45, 1, 0, 39, 39, 1, 0, 48, 57, 2, 0, 65,
		65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100,
		100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103,
		103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106,
		106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109,
		109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112,
		112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115,
		115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118,
		118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121,
		121, 2, 0, 90, 90, 122, 122, 453, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0,
		0, 0, 1, 137, 1, 0, 0, 0, 3, 142, 1, 0, 0, 0, 5, 148, 1, 0, 0, 0, 7, 150,
		1, 0, 0, 0, 9, 152, 1, 0, 0, 0, 11, 154, 1, 0, 0, 0, 13, 158, 1, 0, 0,
		0, 15, 160, 1, 0, 0, 0, 17, 167, 1, 0, 0, 0, 19, 173, 1, 0, 0, 0, 21, 189,
		1, 0, 0, 0, 23, 198, 1, 0, 0, 0, 25, 202, 1, 0, 0, 0, 27, 205, 1, 0, 0,
		0, 29, 208, 1, 0, 0, 0, 31, 213, 1, 0, 0, 0, 33, 222, 1, 0, 0, 0, 35, 224,
		1, 0, 0, 0, 37, 226, 1, 0, 0, 0, 39, 228, 1, 0, 0, 0, 41, 231, 1, 0, 0,
		0, 43, 234, 1, 0, 0, 0, 45, 237, 1, 0, 0, 0, 47, 242, 1, 0, 0, 0, 49, 246,
		1, 0, 0, 0, 51, 252, 1, 0, 0, 0, 53, 256, 1, 0, 0, 0, 55, 261, 1, 0, 0,
		0, 57, 264, 1, 0, 0, 0, 59, 270, 1, 0, 0, 0, 61, 274, 1, 0, 0, 0, 63, 278,
		1, 0, 0, 0, 65, 282, 1, 0, 0, 0, 67, 286, 1, 0, 0, 0, 69, 291, 1, 0, 0,
		0, 71, 297, 1, 0, 0, 0, 73, 303, 1, 0, 0, 0, 75, 344, 1, 0, 0, 0, 77, 388,
		1, 0, 0, 0, 79, 390, 1, 0, 0, 0, 81, 401, 1, 0, 0, 0, 83, 403, 1, 0, 0,
		0, 85, 405, 1, 0, 0, 0, 87, 407, 1, 0, 0, 0, 89, 409, 1, 0, 0, 0, 91, 411,
		1, 0, 0, 0, 93, 413, 1, 0, 0, 0, 95, 415, 1, 0, 0, 0, 97, 417, 1, 0, 0,
		0, 99, 419, 1, 0, 0, 0, 101, 421, 1, 0, 0, 0, 103, 423, 1, 0, 0, 0, 105,
		425, 1, 0, 0, 0, 107, 427, 1, 0, 0, 0, 109, 429, 1, 0, 0, 0, 111, 431,
		1, 0, 0, 0, 113, 433, 1, 0, 0, 0, 115, 435, 1, 0, 0, 0, 117, 437, 1, 0,
		0, 0, 119, 439, 1, 0, 0, 0, 121, 441, 1, 0, 0, 0, 123, 443, 1, 0, 0, 0,
		125, 445, 1, 0, 0, 0, 127, 447, 1, 0, 0, 0, 129, 449, 1, 0, 0, 0, 131,
		451, 1, 0, 0, 0, 133, 453, 1, 0, 0, 0, 135, 455, 1, 0, 0, 0, 137, 138,
		7, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 6, 0, 0, 0, 140, 2, 1, 0, 0,
		0, 141, 143, 7, 1, 0, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147,
		6, 1, 1, 0, 147, 4, 1, 0, 0, 0, 148, 149, 5, 44, 0, 0, 149, 6, 1, 0, 0,
		0, 150, 151, 5, 40, 0, 0, 151, 8, 1, 0, 0, 0, 152, 153, 5, 41, 0, 0, 153,
		10, 1, 0, 0, 0, 154, 155, 5, 59, 0, 0, 155, 12, 1, 0, 0, 0, 156, 159, 3,
		67, 33, 0, 157, 159, 3, 69, 34, 0, 158, 156, 1, 0, 0, 0, 158, 157, 1, 0,
		0, 0, 159, 14, 1, 0, 0, 0, 160, 161, 3, 121, 60, 0, 161, 162, 3, 93, 46,
		0, 162, 163, 3, 107, 53, 0, 163, 164, 3, 93, 46, 0, 164, 165, 3, 89, 44,
		0, 165, 166, 3, 123, 61, 0, 166, 16, 1, 0, 0, 0, 167, 168, 3, 129, 64,
		0, 168, 169, 3, 99, 49, 0, 169, 170, 3, 93, 46, 0, 170, 171, 3, 119, 59,
		0, 171, 172, 3, 93, 46, 0, 172, 18, 1, 0, 0, 0, 173, 174, 3, 129, 64, 0,
		174, 175, 3, 101, 50, 0, 175, 176, 3, 111, 55, 0, 176, 177, 3, 91, 45,
		0, 177, 178, 3, 113, 56, 0, 178, 179, 3, 129, 64, 0, 179, 180, 3, 1, 0,
		0, 180, 181, 3, 123, 61, 0, 181, 182, 3, 125, 62, 0, 182, 183, 3, 109,
		54, 0, 183, 184, 3, 87, 43, 0, 184, 185, 3, 107, 53, 0, 185, 186, 3, 101,
		50, 0, 186, 187, 3, 111, 55, 0, 187, 188, 3, 97, 48, 0, 188, 20, 1, 0,
		0, 0, 189, 190, 3, 97, 48, 0, 190, 191, 3, 119, 59, 0, 191, 192, 3, 113,
		56, 0, 192, 193, 3, 125, 62, 0, 193, 194, 3, 115, 57, 0, 194, 195, 3, 1,
		0, 0, 195, 196, 3, 87, 43, 0, 196, 197, 3, 133, 66, 0, 197, 22, 1, 0, 0,
		0, 198, 199, 3, 85, 42, 0, 199, 200, 3, 111, 55, 0, 200, 201, 3, 91, 45,
		0, 201, 24, 1, 0, 0, 0, 202, 203, 3, 113, 56, 0, 203, 204, 3, 119, 59,
		0, 204, 26, 1, 0, 0, 0, 205, 206, 3, 101, 50, 0, 206, 207, 3, 121, 60,
		0, 207, 28, 1, 0, 0, 0, 208, 209, 3, 107, 53, 0, 209, 210, 3, 101, 50,
		0, 210, 211, 3, 105, 52, 0, 211, 212, 3, 93, 46, 0, 212, 30, 1, 0, 0, 0,
		213, 214, 3, 111, 55, 0, 214, 215, 3, 113, 56, 0, 215, 216, 3, 123, 61,
		0, 216, 217, 3, 1, 0, 0, 217, 218, 3, 107, 53, 0, 218, 219, 3, 101, 50,
		0, 219, 220, 3, 105, 52, 0, 220, 221, 3, 93, 46, 0, 221, 32, 1, 0, 0, 0,
		222, 223, 5, 61, 0, 0, 223, 34, 1, 0, 0, 0, 224, 225, 5, 62, 0, 0, 225,
		36, 1, 0, 0, 0, 226, 227, 5, 60, 0, 0, 227, 38, 1, 0, 0, 0, 228, 229, 3,
		37, 18, 0, 229, 230, 3, 33, 16, 0, 230, 40, 1, 0, 0, 0, 231, 232, 3, 35,
		17, 0, 232, 233, 3, 33, 16, 0, 233, 42, 1, 0, 0, 0, 234, 235, 5, 33, 0,
		0, 235, 236, 3, 33, 16, 0, 236, 44, 1, 0, 0, 0, 237, 238, 3, 111, 55, 0,
		238, 239, 3, 125, 62, 0, 239, 240, 3, 107, 53, 0, 240, 241, 3, 107, 53,
		0, 241, 46, 1, 0, 0, 0, 242, 243, 3, 27, 13, 0, 243, 244, 3, 1, 0, 0, 244,
		245, 3, 45, 22, 0, 245, 48, 1, 0, 0, 0, 246, 247, 3, 27, 13, 0, 247, 248,
		3, 1, 0, 0, 248, 249, 3, 51, 25, 0, 249, 250, 3, 1, 0, 0, 250, 251, 3,
		45, 22, 0, 251, 50, 1, 0, 0, 0, 252, 253, 3, 111, 55, 0, 253, 254, 3, 113,
		56, 0, 254, 255, 3, 123, 61, 0, 255, 52, 1, 0, 0, 0, 256, 257, 3, 51, 25,
		0, 257, 258, 3, 1, 0, 0, 258, 259, 3, 101, 50, 0, 259, 260, 3, 111, 55,
		0, 260, 54, 1, 0, 0, 0, 261, 262, 3, 101, 50, 0, 262, 263, 3, 111, 55,
		0, 263, 56, 1, 0, 0, 0, 264, 265, 3, 89, 44, 0, 265, 266, 3, 113, 56, 0,
		266, 267, 3, 125, 62, 0, 267, 268, 3, 111, 55, 0, 268, 269, 3, 123, 61,
		0, 269, 58, 1, 0, 0, 0, 270, 271, 3, 121, 60, 0, 271, 272, 3, 125, 62,
		0, 272, 273, 3, 109, 54, 0, 273, 60, 1, 0, 0, 0, 274, 275, 3, 109, 54,
		0, 275, 276, 3, 101, 50, 0, 276, 277, 3, 111, 55, 0, 277, 62, 1, 0, 0,
		0, 278, 279, 3, 109, 54, 0, 279, 280, 3, 85, 42, 0, 280, 281, 3, 131, 65,
		0, 281, 64, 1, 0, 0, 0, 282, 283, 3, 85, 42, 0, 283, 284, 3, 127, 63, 0,
		284, 285, 3, 97, 48, 0, 285, 66, 1, 0, 0, 0, 286, 287, 3, 123, 61, 0, 287,
		288, 3, 119, 59, 0, 288, 289, 3, 125, 62, 0, 289, 290, 3, 93, 46, 0, 290,
		68, 1, 0, 0, 0, 291, 292, 3, 95, 47, 0, 292, 293, 3, 85, 42, 0, 293, 294,
		3, 107, 53, 0, 294, 295, 3, 121, 60, 0, 295, 296, 3, 93, 46, 0, 296, 70,
		1, 0, 0, 0, 297, 298, 3, 125, 62, 0, 298, 299, 3, 115, 57, 0, 299, 300,
		3, 115, 57, 0, 300, 301, 3, 93, 46, 0, 301, 302, 3, 119, 59, 0, 302, 72,
		1, 0, 0, 0, 303, 304, 3, 107, 53, 0, 304, 305, 3, 113, 56, 0, 305, 306,
		3, 129, 64, 0, 306, 307, 3, 93, 46, 0, 307, 308, 3, 119, 59, 0, 308, 74,
		1, 0, 0, 0, 309, 315, 5, 34, 0, 0, 310, 314, 8, 2, 0, 0, 311, 312, 5, 34,
		0, 0, 312, 314, 5, 34, 0, 0, 313, 310, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0,
		314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316,
		318, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 345, 5, 34, 0, 0, 319, 325,
		5, 96, 0, 0, 320, 324, 8, 3, 0, 0, 321, 322, 5, 96, 0, 0, 322, 324, 5,
		96, 0, 0, 323, 320, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 327, 1, 0, 0,
		0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 328, 1, 0, 0, 0, 327,
		325, 1, 0, 0, 0, 328, 345, 5, 96, 0, 0, 329, 333, 5, 91, 0, 0, 330, 332,
		8, 4, 0, 0, 331, 330, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0,
		0, 0, 333, 334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0,
		336, 345, 5, 93, 0, 0, 337, 341, 7, 5, 0, 0, 338, 340, 7, 6, 0, 0, 339,
		338, 1, 0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342,
		1, 0, 0, 0, 342, 345, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 309, 1, 0,
		0, 0, 344, 319, 1, 0, 0, 0, 344, 329, 1, 0, 0, 0, 344, 337, 1, 0, 0, 0,
		345, 76, 1, 0, 0, 0, 346, 348, 3, 83, 41, 0, 347, 346, 1, 0, 0, 0, 348,
		349, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 358,
		1, 0, 0, 0, 351, 355, 5, 46, 0, 0, 352, 354, 3, 83, 41, 0, 353, 352, 1,
		0, 0, 0, 354, 357, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0,
		0, 356, 359, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 358, 351, 1, 0, 0, 0, 358,
		359, 1, 0, 0, 0, 359, 369, 1, 0, 0, 0, 360, 362, 3, 93, 46, 0, 361, 363,
		7, 7, 0, 0, 362, 361, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 365, 1, 0,
		0, 0, 364, 366, 3, 83, 41, 0, 365, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0,
		0, 367, 365, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 370, 1, 0, 0, 0, 369,
		360, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 389, 1, 0, 0, 0, 371, 373,
		5, 46, 0, 0, 372, 374, 3, 83, 41, 0, 373, 372, 1, 0, 0, 0, 374, 375, 1,
		0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 386, 1, 0, 0,
		0, 377, 379, 3, 93, 46, 0, 378, 380, 7, 7, 0, 0, 379, 378, 1, 0, 0, 0,
		379, 380, 1, 0, 0, 0, 380, 382, 1, 0, 0, 0, 381, 383, 3, 83, 41, 0, 382,
		381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385,
		1, 0, 0, 0, 385, 387, 1, 0, 0, 0, 386, 377, 1, 0, 0, 0, 386, 387, 1, 0,
		0, 0, 387, 389, 1, 0, 0, 0, 388, 347, 1, 0, 0, 0, 388, 371, 1, 0, 0, 0,
		389, 78, 1, 0, 0, 0, 390, 396, 5, 39, 0, 0, 391, 395, 8, 8, 0, 0, 392,
		393, 5, 39, 0, 0, 393, 395, 5, 39, 0, 0, 394, 391, 1, 0, 0, 0, 394, 392,
		1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397, 1, 0,
		0, 0, 397, 399, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 400, 5, 39, 0, 0,
		400, 80, 1, 0, 0, 0, 401, 402, 5, 42, 0, 0, 402, 82, 1, 0, 0, 0, 403, 404,
		7, 9, 0, 0, 404, 84, 1, 0, 0, 0, 405, 406, 7, 10, 0, 0, 406, 86, 1, 0,
		0, 0, 407, 408, 7, 11, 0, 0, 408, 88, 1, 0, 0, 0, 409, 410, 7, 12, 0, 0,
		410, 90, 1, 0, 0, 0, 411, 412, 7, 13, 0, 0, 412, 92, 1, 0, 0, 0, 413, 414,
		7, 14, 0, 0, 414, 94, 1, 0, 0, 0, 415, 416, 7, 15, 0, 0, 416, 96, 1, 0,
		0, 0, 417, 418, 7, 16, 0, 0, 418, 98, 1, 0, 0, 0, 419, 420, 7, 17, 0, 0,
		420, 100, 1, 0, 0, 0, 421, 422, 7, 18, 0, 0, 422, 102, 1, 0, 0, 0, 423,
		424, 7, 19, 0, 0, 424, 104, 1, 0, 0, 0, 425, 426, 7, 20, 0, 0, 426, 106,
		1, 0, 0, 0, 427, 428, 7, 21, 0, 0, 428, 108, 1, 0, 0, 0, 429, 430, 7, 22,
		0, 0, 430, 110, 1, 0, 0, 0, 431, 432, 7, 23, 0, 0, 432, 112, 1, 0, 0, 0,
		433, 434, 7, 24, 0, 0, 434, 114, 1, 0, 0, 0, 435, 436, 7, 25, 0, 0, 436,
		116, 1, 0, 0, 0, 437, 438, 7, 26, 0, 0, 438, 118, 1, 0, 0, 0, 439, 440,
		7, 27, 0, 0, 440, 120, 1, 0, 0, 0, 441, 442, 7, 28, 0, 0, 442, 122, 1,
		0, 0, 0, 443, 444, 7, 29, 0, 0, 444, 124, 1, 0, 0, 0, 445, 446, 7, 30,
		0, 0, 446, 126, 1, 0, 0, 0, 447, 448, 7, 31, 0, 0, 448, 128, 1, 0, 0, 0,
		449, 450, 7, 32, 0, 0, 450, 130, 1, 0, 0, 0, 451, 452, 7, 33, 0, 0, 452,
		132, 1, 0, 0, 0, 453, 454, 7, 34, 0, 0, 454, 134, 1, 0, 0, 0, 455, 456,
		7, 35, 0, 0, 456, 136, 1, 0, 0, 0, 23, 0, 144, 158, 313, 315, 323, 325,
		333, 341, 344, 349, 355, 358, 362, 367, 369, 375, 379, 384, 386, 388, 394,
		396, 2, 0, 1, 0, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SqlLexerInit initializes any static state used to implement SqlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSqlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SqlLexerInit() {
	staticData := &sqllexerLexerStaticData
	staticData.once.Do(sqllexerLexerInit)
}

// NewSqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSqlLexer(input antlr.CharStream) *SqlLexer {
	SqlLexerInit()
	l := new(SqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &sqllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Sql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SqlLexer tokens.
const (
	SqlLexerSPACE             = 1
	SqlLexerWS                = 2
	SqlLexerCOMMA             = 3
	SqlLexerL_BRACKET         = 4
	SqlLexerR_BRACKET         = 5
	SqlLexerEOQ               = 6
	SqlLexerBOOLEAN_LITERAL   = 7
	SqlLexerK_SELECT          = 8
	SqlLexerK_WHERE           = 9
	SqlLexerK_WINDOW_TUMBLING = 10
	SqlLexerK_GROUP_BY        = 11
	SqlLexerK_AND             = 12
	SqlLexerK_OR              = 13
	SqlLexerK_IS              = 14
	SqlLexerK_LIKE            = 15
	SqlLexerK_NOT_LIKE        = 16
	SqlLexerK_EQUAL           = 17
	SqlLexerK_GREATER         = 18
	SqlLexerK_LESS            = 19
	SqlLexerK_LESS_EQUAL      = 20
	SqlLexerK_GREATER_EQUAL   = 21
	SqlLexerK_NOT_EQUAL       = 22
	SqlLexerK_NULL            = 23
	SqlLexerK_IS_NULL         = 24
	SqlLexerK_IS_NOT_NULL     = 25
	SqlLexerK_NOT             = 26
	SqlLexerK_NOT_IN          = 27
	SqlLexerK_IN              = 28
	SqlLexerK_COUNT           = 29
	SqlLexerK_SUM             = 30
	SqlLexerK_MIN             = 31
	SqlLexerK_MAX             = 32
	SqlLexerK_AVG             = 33
	SqlLexerK_TRUE            = 34
	SqlLexerK_FALSE           = 35
	SqlLexerK_UPPER           = 36
	SqlLexerK_LOWER           = 37
	SqlLexerIDENTIFIER        = 38
	SqlLexerNUMERIC_LITERAL   = 39
	SqlLexerSTRING_LITERAL    = 40
	SqlLexerSTAR              = 41
)
