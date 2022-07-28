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
		"", "'='", "'>'", "'<'", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "'*'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "WS", "COMMA", "L_BRACKET", "R_BRACKET", "EOQ", "K_SELECT",
		"K_WHERE", "K_WINDOW_TUMBLING", "K_GROUP_BY", "K_AND", "K_OR", "K_IS",
		"K_LIKE", "K_EQUAL", "K_GREATER", "K_LESS", "K_LESS_EQUAL", "K_GREATER_EQUAL",
		"K_NOT_EQUAL", "K_NULL", "K_IS_NULL", "K_IS_NOT_NULL", "K_NOT", "K_NOT_IN",
		"K_IN", "K_TRUE", "K_FALSE", "K_COUNT", "K_MIN", "K_MAX", "K_AVG", "IDENTIFIER",
		"BOOLEAN_LITERAL", "NUMERIC_LITERAL", "STRING_LITERAL", "STAR",
	}
	staticData.ruleNames = []string{
		"SPACE", "WS", "COMMA", "L_BRACKET", "R_BRACKET", "EOQ", "K_SELECT",
		"K_WHERE", "K_WINDOW_TUMBLING", "K_GROUP_BY", "K_AND", "K_OR", "K_IS",
		"K_LIKE", "K_EQUAL", "K_GREATER", "K_LESS", "K_LESS_EQUAL", "K_GREATER_EQUAL",
		"K_NOT_EQUAL", "K_NULL", "K_IS_NULL", "K_IS_NOT_NULL", "K_NOT", "K_NOT_IN",
		"K_IN", "K_TRUE", "K_FALSE", "K_COUNT", "K_MIN", "K_MAX", "K_AVG", "IDENTIFIER",
		"BOOLEAN_LITERAL", "NUMERIC_LITERAL", "STRING_LITERAL", "STAR", "DIGIT",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 424, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		62, 2, 63, 7, 63, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 135, 8, 1, 11, 1,
		12, 1, 136, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 277, 8,
		32, 10, 32, 12, 32, 280, 9, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32,
		287, 8, 32, 10, 32, 12, 32, 290, 9, 32, 1, 32, 1, 32, 1, 32, 5, 32, 295,
		8, 32, 10, 32, 12, 32, 298, 9, 32, 1, 32, 1, 32, 1, 32, 5, 32, 303, 8,
		32, 10, 32, 12, 32, 306, 9, 32, 3, 32, 308, 8, 32, 1, 33, 1, 33, 3, 33,
		312, 8, 33, 1, 34, 4, 34, 315, 8, 34, 11, 34, 12, 34, 316, 1, 34, 1, 34,
		5, 34, 321, 8, 34, 10, 34, 12, 34, 324, 9, 34, 3, 34, 326, 8, 34, 1, 34,
		1, 34, 3, 34, 330, 8, 34, 1, 34, 4, 34, 333, 8, 34, 11, 34, 12, 34, 334,
		3, 34, 337, 8, 34, 1, 34, 1, 34, 4, 34, 341, 8, 34, 11, 34, 12, 34, 342,
		1, 34, 1, 34, 3, 34, 347, 8, 34, 1, 34, 4, 34, 350, 8, 34, 11, 34, 12,
		34, 351, 3, 34, 354, 8, 34, 3, 34, 356, 8, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 5, 35, 362, 8, 35, 10, 35, 12, 35, 365, 9, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46,
		1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1,
		52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57,
		1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1,
		62, 1, 63, 1, 63, 0, 0, 64, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69,
		35, 71, 36, 73, 37, 75, 0, 77, 0, 79, 0, 81, 0, 83, 0, 85, 0, 87, 0, 89,
		0, 91, 0, 93, 0, 95, 0, 97, 0, 99, 0, 101, 0, 103, 0, 105, 0, 107, 0, 109,
		0, 111, 0, 113, 0, 115, 0, 117, 0, 119, 0, 121, 0, 123, 0, 125, 0, 127,
		0, 1, 0, 36, 3, 0, 9, 11, 13, 13, 32, 32, 2, 0, 9, 9, 32, 32, 1, 0, 34,
		34, 1, 0, 96, 96, 1, 0, 93, 93, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 2, 0, 43, 43, 45, 45, 1, 0, 39, 39, 1, 0,
		48, 57, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99,
		2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102,
		2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105,
		2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108,
		2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111,
		2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114,
		2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117,
		2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120,
		2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 420, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 1, 129, 1, 0, 0, 0, 3, 134, 1, 0, 0, 0, 5, 140,
		1, 0, 0, 0, 7, 142, 1, 0, 0, 0, 9, 144, 1, 0, 0, 0, 11, 146, 1, 0, 0, 0,
		13, 148, 1, 0, 0, 0, 15, 155, 1, 0, 0, 0, 17, 161, 1, 0, 0, 0, 19, 177,
		1, 0, 0, 0, 21, 186, 1, 0, 0, 0, 23, 190, 1, 0, 0, 0, 25, 193, 1, 0, 0,
		0, 27, 196, 1, 0, 0, 0, 29, 201, 1, 0, 0, 0, 31, 203, 1, 0, 0, 0, 33, 205,
		1, 0, 0, 0, 35, 207, 1, 0, 0, 0, 37, 210, 1, 0, 0, 0, 39, 213, 1, 0, 0,
		0, 41, 216, 1, 0, 0, 0, 43, 221, 1, 0, 0, 0, 45, 225, 1, 0, 0, 0, 47, 231,
		1, 0, 0, 0, 49, 235, 1, 0, 0, 0, 51, 240, 1, 0, 0, 0, 53, 243, 1, 0, 0,
		0, 55, 248, 1, 0, 0, 0, 57, 254, 1, 0, 0, 0, 59, 260, 1, 0, 0, 0, 61, 264,
		1, 0, 0, 0, 63, 268, 1, 0, 0, 0, 65, 307, 1, 0, 0, 0, 67, 311, 1, 0, 0,
		0, 69, 355, 1, 0, 0, 0, 71, 357, 1, 0, 0, 0, 73, 368, 1, 0, 0, 0, 75, 370,
		1, 0, 0, 0, 77, 372, 1, 0, 0, 0, 79, 374, 1, 0, 0, 0, 81, 376, 1, 0, 0,
		0, 83, 378, 1, 0, 0, 0, 85, 380, 1, 0, 0, 0, 87, 382, 1, 0, 0, 0, 89, 384,
		1, 0, 0, 0, 91, 386, 1, 0, 0, 0, 93, 388, 1, 0, 0, 0, 95, 390, 1, 0, 0,
		0, 97, 392, 1, 0, 0, 0, 99, 394, 1, 0, 0, 0, 101, 396, 1, 0, 0, 0, 103,
		398, 1, 0, 0, 0, 105, 400, 1, 0, 0, 0, 107, 402, 1, 0, 0, 0, 109, 404,
		1, 0, 0, 0, 111, 406, 1, 0, 0, 0, 113, 408, 1, 0, 0, 0, 115, 410, 1, 0,
		0, 0, 117, 412, 1, 0, 0, 0, 119, 414, 1, 0, 0, 0, 121, 416, 1, 0, 0, 0,
		123, 418, 1, 0, 0, 0, 125, 420, 1, 0, 0, 0, 127, 422, 1, 0, 0, 0, 129,
		130, 7, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 6, 0, 0, 0, 132, 2, 1,
		0, 0, 0, 133, 135, 7, 1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0,
		0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138,
		139, 6, 1, 1, 0, 139, 4, 1, 0, 0, 0, 140, 141, 5, 44, 0, 0, 141, 6, 1,
		0, 0, 0, 142, 143, 5, 40, 0, 0, 143, 8, 1, 0, 0, 0, 144, 145, 5, 41, 0,
		0, 145, 10, 1, 0, 0, 0, 146, 147, 5, 59, 0, 0, 147, 12, 1, 0, 0, 0, 148,
		149, 3, 113, 56, 0, 149, 150, 3, 85, 42, 0, 150, 151, 3, 99, 49, 0, 151,
		152, 3, 85, 42, 0, 152, 153, 3, 81, 40, 0, 153, 154, 3, 115, 57, 0, 154,
		14, 1, 0, 0, 0, 155, 156, 3, 121, 60, 0, 156, 157, 3, 91, 45, 0, 157, 158,
		3, 85, 42, 0, 158, 159, 3, 111, 55, 0, 159, 160, 3, 85, 42, 0, 160, 16,
		1, 0, 0, 0, 161, 162, 3, 121, 60, 0, 162, 163, 3, 93, 46, 0, 163, 164,
		3, 103, 51, 0, 164, 165, 3, 83, 41, 0, 165, 166, 3, 105, 52, 0, 166, 167,
		3, 121, 60, 0, 167, 168, 3, 1, 0, 0, 168, 169, 3, 115, 57, 0, 169, 170,
		3, 117, 58, 0, 170, 171, 3, 101, 50, 0, 171, 172, 3, 79, 39, 0, 172, 173,
		3, 99, 49, 0, 173, 174, 3, 93, 46, 0, 174, 175, 3, 103, 51, 0, 175, 176,
		3, 89, 44, 0, 176, 18, 1, 0, 0, 0, 177, 178, 3, 89, 44, 0, 178, 179, 3,
		111, 55, 0, 179, 180, 3, 105, 52, 0, 180, 181, 3, 117, 58, 0, 181, 182,
		3, 107, 53, 0, 182, 183, 3, 1, 0, 0, 183, 184, 3, 79, 39, 0, 184, 185,
		3, 125, 62, 0, 185, 20, 1, 0, 0, 0, 186, 187, 3, 77, 38, 0, 187, 188, 3,
		103, 51, 0, 188, 189, 3, 83, 41, 0, 189, 22, 1, 0, 0, 0, 190, 191, 3, 105,
		52, 0, 191, 192, 3, 111, 55, 0, 192, 24, 1, 0, 0, 0, 193, 194, 3, 93, 46,
		0, 194, 195, 3, 113, 56, 0, 195, 26, 1, 0, 0, 0, 196, 197, 3, 99, 49, 0,
		197, 198, 3, 93, 46, 0, 198, 199, 3, 97, 48, 0, 199, 200, 3, 85, 42, 0,
		200, 28, 1, 0, 0, 0, 201, 202, 5, 61, 0, 0, 202, 30, 1, 0, 0, 0, 203, 204,
		5, 62, 0, 0, 204, 32, 1, 0, 0, 0, 205, 206, 5, 60, 0, 0, 206, 34, 1, 0,
		0, 0, 207, 208, 3, 33, 16, 0, 208, 209, 3, 29, 14, 0, 209, 36, 1, 0, 0,
		0, 210, 211, 3, 31, 15, 0, 211, 212, 3, 29, 14, 0, 212, 38, 1, 0, 0, 0,
		213, 214, 5, 33, 0, 0, 214, 215, 3, 29, 14, 0, 215, 40, 1, 0, 0, 0, 216,
		217, 3, 103, 51, 0, 217, 218, 3, 117, 58, 0, 218, 219, 3, 99, 49, 0, 219,
		220, 3, 99, 49, 0, 220, 42, 1, 0, 0, 0, 221, 222, 3, 25, 12, 0, 222, 223,
		3, 1, 0, 0, 223, 224, 3, 41, 20, 0, 224, 44, 1, 0, 0, 0, 225, 226, 3, 25,
		12, 0, 226, 227, 3, 1, 0, 0, 227, 228, 3, 47, 23, 0, 228, 229, 3, 1, 0,
		0, 229, 230, 3, 41, 20, 0, 230, 46, 1, 0, 0, 0, 231, 232, 3, 103, 51, 0,
		232, 233, 3, 105, 52, 0, 233, 234, 3, 115, 57, 0, 234, 48, 1, 0, 0, 0,
		235, 236, 3, 47, 23, 0, 236, 237, 3, 1, 0, 0, 237, 238, 3, 93, 46, 0, 238,
		239, 3, 103, 51, 0, 239, 50, 1, 0, 0, 0, 240, 241, 3, 93, 46, 0, 241, 242,
		3, 103, 51, 0, 242, 52, 1, 0, 0, 0, 243, 244, 3, 115, 57, 0, 244, 245,
		3, 111, 55, 0, 245, 246, 3, 117, 58, 0, 246, 247, 3, 85, 42, 0, 247, 54,
		1, 0, 0, 0, 248, 249, 3, 87, 43, 0, 249, 250, 3, 77, 38, 0, 250, 251, 3,
		99, 49, 0, 251, 252, 3, 113, 56, 0, 252, 253, 3, 85, 42, 0, 253, 56, 1,
		0, 0, 0, 254, 255, 3, 81, 40, 0, 255, 256, 3, 105, 52, 0, 256, 257, 3,
		117, 58, 0, 257, 258, 3, 103, 51, 0, 258, 259, 3, 115, 57, 0, 259, 58,
		1, 0, 0, 0, 260, 261, 3, 101, 50, 0, 261, 262, 3, 93, 46, 0, 262, 263,
		3, 103, 51, 0, 263, 60, 1, 0, 0, 0, 264, 265, 3, 101, 50, 0, 265, 266,
		3, 77, 38, 0, 266, 267, 3, 123, 61, 0, 267, 62, 1, 0, 0, 0, 268, 269, 3,
		77, 38, 0, 269, 270, 3, 119, 59, 0, 270, 271, 3, 89, 44, 0, 271, 64, 1,
		0, 0, 0, 272, 278, 5, 34, 0, 0, 273, 277, 8, 2, 0, 0, 274, 275, 5, 34,
		0, 0, 275, 277, 5, 34, 0, 0, 276, 273, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0,
		277, 280, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279,
		281, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 308, 5, 34, 0, 0, 282, 288,
		5, 96, 0, 0, 283, 287, 8, 3, 0, 0, 284, 285, 5, 96, 0, 0, 285, 287, 5,
		96, 0, 0, 286, 283, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 287, 290, 1, 0, 0,
		0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1, 0, 0, 0, 290,
		288, 1, 0, 0, 0, 291, 308, 5, 96, 0, 0, 292, 296, 5, 91, 0, 0, 293, 295,
		8, 4, 0, 0, 294, 293, 1, 0, 0, 0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0,
		0, 0, 296, 297, 1, 0, 0, 0, 297, 299, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0,
		299, 308, 5, 93, 0, 0, 300, 304, 7, 5, 0, 0, 301, 303, 7, 6, 0, 0, 302,
		301, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305,
		1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 272, 1, 0,
		0, 0, 307, 282, 1, 0, 0, 0, 307, 292, 1, 0, 0, 0, 307, 300, 1, 0, 0, 0,
		308, 66, 1, 0, 0, 0, 309, 312, 3, 53, 26, 0, 310, 312, 3, 55, 27, 0, 311,
		309, 1, 0, 0, 0, 311, 310, 1, 0, 0, 0, 312, 68, 1, 0, 0, 0, 313, 315, 3,
		75, 37, 0, 314, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 314, 1, 0,
		0, 0, 316, 317, 1, 0, 0, 0, 317, 325, 1, 0, 0, 0, 318, 322, 5, 46, 0, 0,
		319, 321, 3, 75, 37, 0, 320, 319, 1, 0, 0, 0, 321, 324, 1, 0, 0, 0, 322,
		320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322,
		1, 0, 0, 0, 325, 318, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 336, 1, 0,
		0, 0, 327, 329, 3, 85, 42, 0, 328, 330, 7, 7, 0, 0, 329, 328, 1, 0, 0,
		0, 329, 330, 1, 0, 0, 0, 330, 332, 1, 0, 0, 0, 331, 333, 3, 75, 37, 0,
		332, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 334,
		335, 1, 0, 0, 0, 335, 337, 1, 0, 0, 0, 336, 327, 1, 0, 0, 0, 336, 337,
		1, 0, 0, 0, 337, 356, 1, 0, 0, 0, 338, 340, 5, 46, 0, 0, 339, 341, 3, 75,
		37, 0, 340, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0,
		342, 343, 1, 0, 0, 0, 343, 353, 1, 0, 0, 0, 344, 346, 3, 85, 42, 0, 345,
		347, 7, 7, 0, 0, 346, 345, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 349,
		1, 0, 0, 0, 348, 350, 3, 75, 37, 0, 349, 348, 1, 0, 0, 0, 350, 351, 1,
		0, 0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 354, 1, 0, 0,
		0, 353, 344, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 356, 1, 0, 0, 0, 355,
		314, 1, 0, 0, 0, 355, 338, 1, 0, 0, 0, 356, 70, 1, 0, 0, 0, 357, 363, 5,
		39, 0, 0, 358, 362, 8, 8, 0, 0, 359, 360, 5, 39, 0, 0, 360, 362, 5, 39,
		0, 0, 361, 358, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0,
		363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 366, 1, 0, 0, 0, 365,
		363, 1, 0, 0, 0, 366, 367, 5, 39, 0, 0, 367, 72, 1, 0, 0, 0, 368, 369,
		5, 42, 0, 0, 369, 74, 1, 0, 0, 0, 370, 371, 7, 9, 0, 0, 371, 76, 1, 0,
		0, 0, 372, 373, 7, 10, 0, 0, 373, 78, 1, 0, 0, 0, 374, 375, 7, 11, 0, 0,
		375, 80, 1, 0, 0, 0, 376, 377, 7, 12, 0, 0, 377, 82, 1, 0, 0, 0, 378, 379,
		7, 13, 0, 0, 379, 84, 1, 0, 0, 0, 380, 381, 7, 14, 0, 0, 381, 86, 1, 0,
		0, 0, 382, 383, 7, 15, 0, 0, 383, 88, 1, 0, 0, 0, 384, 385, 7, 16, 0, 0,
		385, 90, 1, 0, 0, 0, 386, 387, 7, 17, 0, 0, 387, 92, 1, 0, 0, 0, 388, 389,
		7, 18, 0, 0, 389, 94, 1, 0, 0, 0, 390, 391, 7, 19, 0, 0, 391, 96, 1, 0,
		0, 0, 392, 393, 7, 20, 0, 0, 393, 98, 1, 0, 0, 0, 394, 395, 7, 21, 0, 0,
		395, 100, 1, 0, 0, 0, 396, 397, 7, 22, 0, 0, 397, 102, 1, 0, 0, 0, 398,
		399, 7, 23, 0, 0, 399, 104, 1, 0, 0, 0, 400, 401, 7, 24, 0, 0, 401, 106,
		1, 0, 0, 0, 402, 403, 7, 25, 0, 0, 403, 108, 1, 0, 0, 0, 404, 405, 7, 26,
		0, 0, 405, 110, 1, 0, 0, 0, 406, 407, 7, 27, 0, 0, 407, 112, 1, 0, 0, 0,
		408, 409, 7, 28, 0, 0, 409, 114, 1, 0, 0, 0, 410, 411, 7, 29, 0, 0, 411,
		116, 1, 0, 0, 0, 412, 413, 7, 30, 0, 0, 413, 118, 1, 0, 0, 0, 414, 415,
		7, 31, 0, 0, 415, 120, 1, 0, 0, 0, 416, 417, 7, 32, 0, 0, 417, 122, 1,
		0, 0, 0, 418, 419, 7, 33, 0, 0, 419, 124, 1, 0, 0, 0, 420, 421, 7, 34,
		0, 0, 421, 126, 1, 0, 0, 0, 422, 423, 7, 35, 0, 0, 423, 128, 1, 0, 0, 0,
		23, 0, 136, 276, 278, 286, 288, 296, 304, 307, 311, 316, 322, 325, 329,
		334, 336, 342, 346, 351, 353, 355, 361, 363, 2, 0, 1, 0, 6, 0, 0,
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
	SqlLexerK_SELECT          = 7
	SqlLexerK_WHERE           = 8
	SqlLexerK_WINDOW_TUMBLING = 9
	SqlLexerK_GROUP_BY        = 10
	SqlLexerK_AND             = 11
	SqlLexerK_OR              = 12
	SqlLexerK_IS              = 13
	SqlLexerK_LIKE            = 14
	SqlLexerK_EQUAL           = 15
	SqlLexerK_GREATER         = 16
	SqlLexerK_LESS            = 17
	SqlLexerK_LESS_EQUAL      = 18
	SqlLexerK_GREATER_EQUAL   = 19
	SqlLexerK_NOT_EQUAL       = 20
	SqlLexerK_NULL            = 21
	SqlLexerK_IS_NULL         = 22
	SqlLexerK_IS_NOT_NULL     = 23
	SqlLexerK_NOT             = 24
	SqlLexerK_NOT_IN          = 25
	SqlLexerK_IN              = 26
	SqlLexerK_TRUE            = 27
	SqlLexerK_FALSE           = 28
	SqlLexerK_COUNT           = 29
	SqlLexerK_MIN             = 30
	SqlLexerK_MAX             = 31
	SqlLexerK_AVG             = 32
	SqlLexerIDENTIFIER        = 33
	SqlLexerBOOLEAN_LITERAL   = 34
	SqlLexerNUMERIC_LITERAL   = 35
	SqlLexerSTRING_LITERAL    = 36
	SqlLexerSTAR              = 37
)
