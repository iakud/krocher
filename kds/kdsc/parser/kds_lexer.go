// Code generated from kds.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type kdsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var KdsLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kdslexerLexerInit() {
	staticData := &KdsLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'import'", "'package'", "'repeated'", "'map'", "'int32'", "'int64'",
		"'uint32'", "'uint64'", "'sint32'", "'sint64'", "'fixed32'", "'fixed64'",
		"'sfixed32'", "'sfixed64'", "'bool'", "'string'", "'double'", "'float'",
		"'bytes'", "'Timestamp'", "'Duration'", "'enum'", "'entity'", "'component'",
		"';'", "'='", "'('", "')'", "'['", "']'", "'{'", "'}'", "'<'", "'>'",
		"'.'", "','", "':'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "IMPORT", "PACKAGE", "REPEATED", "MAP", "INT32", "INT64", "UINT32",
		"UINT64", "SINT32", "SINT64", "FIXED32", "FIXED64", "SFIXED32", "SFIXED64",
		"BOOL", "STRING", "DOUBLE", "FLOAT", "BYTES", "TIMESTAMP", "DURATION",
		"ENUM", "ENTITY", "COMPONENT", "SEMI", "EQ", "LP", "RP", "LB", "RB",
		"LC", "RC", "LT", "GT", "DOT", "COMMA", "COLON", "PLUS", "MINUS", "STR_LIT",
		"BOOL_LIT", "INT_LIT", "IDENTIFIER", "WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.RuleNames = []string{
		"IMPORT", "PACKAGE", "REPEATED", "MAP", "INT32", "INT64", "UINT32",
		"UINT64", "SINT32", "SINT64", "FIXED32", "FIXED64", "SFIXED32", "SFIXED64",
		"BOOL", "STRING", "DOUBLE", "FLOAT", "BYTES", "TIMESTAMP", "DURATION",
		"ENUM", "ENTITY", "COMPONENT", "SEMI", "EQ", "LP", "RP", "LB", "RB",
		"LC", "RC", "LT", "GT", "DOT", "COMMA", "COLON", "PLUS", "MINUS", "STR_LIT",
		"CHAR_VALUE", "HEX_ESCAPE", "OCT_ESCAPE", "CHAR_ESCAPE", "BOOL_LIT",
		"INT_LIT", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "IDENTIFIER", "LETTER",
		"DECIMAL_DIGIT", "OCTAL_DIGIT", "HEX_DIGIT", "WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 46, 441, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 322, 8, 39,
		10, 39, 12, 39, 325, 9, 39, 1, 39, 1, 39, 1, 39, 5, 39, 330, 8, 39, 10,
		39, 12, 39, 333, 9, 39, 1, 39, 3, 39, 336, 8, 39, 1, 40, 1, 40, 1, 40,
		1, 40, 3, 40, 342, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 366, 8, 44, 1, 45, 1, 45, 1,
		45, 3, 45, 371, 8, 45, 1, 46, 1, 46, 5, 46, 375, 8, 46, 10, 46, 12, 46,
		378, 9, 46, 1, 47, 1, 47, 5, 47, 382, 8, 47, 10, 47, 12, 47, 385, 9, 47,
		1, 48, 1, 48, 1, 48, 4, 48, 390, 8, 48, 11, 48, 12, 48, 391, 1, 49, 1,
		49, 1, 49, 5, 49, 397, 8, 49, 10, 49, 12, 49, 400, 9, 49, 1, 50, 1, 50,
		1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 4, 54, 411, 8, 54, 11,
		54, 12, 54, 412, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 5, 55, 421,
		8, 55, 10, 55, 12, 55, 424, 9, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1,
		56, 5, 56, 432, 8, 56, 10, 56, 12, 56, 435, 9, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 3, 323, 331, 433, 0, 57, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 0, 83, 0, 85,
		0, 87, 0, 89, 41, 91, 42, 93, 0, 95, 0, 97, 0, 99, 43, 101, 0, 103, 0,
		105, 0, 107, 0, 109, 44, 111, 45, 113, 46, 1, 0, 10, 3, 0, 0, 0, 10, 10,
		92, 92, 2, 0, 88, 88, 120, 120, 9, 0, 34, 34, 39, 39, 92, 92, 97, 98, 102,
		102, 110, 110, 114, 114, 116, 116, 118, 118, 1, 0, 49, 57, 3, 0, 65, 90,
		95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97,
		102, 3, 0, 9, 10, 12, 13, 32, 32, 2, 0, 10, 10, 13, 13, 446, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0,
		0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 1, 115, 1,
		0, 0, 0, 3, 122, 1, 0, 0, 0, 5, 130, 1, 0, 0, 0, 7, 139, 1, 0, 0, 0, 9,
		143, 1, 0, 0, 0, 11, 149, 1, 0, 0, 0, 13, 155, 1, 0, 0, 0, 15, 162, 1,
		0, 0, 0, 17, 169, 1, 0, 0, 0, 19, 176, 1, 0, 0, 0, 21, 183, 1, 0, 0, 0,
		23, 191, 1, 0, 0, 0, 25, 199, 1, 0, 0, 0, 27, 208, 1, 0, 0, 0, 29, 217,
		1, 0, 0, 0, 31, 222, 1, 0, 0, 0, 33, 229, 1, 0, 0, 0, 35, 236, 1, 0, 0,
		0, 37, 242, 1, 0, 0, 0, 39, 248, 1, 0, 0, 0, 41, 258, 1, 0, 0, 0, 43, 267,
		1, 0, 0, 0, 45, 272, 1, 0, 0, 0, 47, 279, 1, 0, 0, 0, 49, 289, 1, 0, 0,
		0, 51, 291, 1, 0, 0, 0, 53, 293, 1, 0, 0, 0, 55, 295, 1, 0, 0, 0, 57, 297,
		1, 0, 0, 0, 59, 299, 1, 0, 0, 0, 61, 301, 1, 0, 0, 0, 63, 303, 1, 0, 0,
		0, 65, 305, 1, 0, 0, 0, 67, 307, 1, 0, 0, 0, 69, 309, 1, 0, 0, 0, 71, 311,
		1, 0, 0, 0, 73, 313, 1, 0, 0, 0, 75, 315, 1, 0, 0, 0, 77, 317, 1, 0, 0,
		0, 79, 335, 1, 0, 0, 0, 81, 341, 1, 0, 0, 0, 83, 343, 1, 0, 0, 0, 85, 348,
		1, 0, 0, 0, 87, 353, 1, 0, 0, 0, 89, 365, 1, 0, 0, 0, 91, 370, 1, 0, 0,
		0, 93, 372, 1, 0, 0, 0, 95, 379, 1, 0, 0, 0, 97, 386, 1, 0, 0, 0, 99, 393,
		1, 0, 0, 0, 101, 401, 1, 0, 0, 0, 103, 403, 1, 0, 0, 0, 105, 405, 1, 0,
		0, 0, 107, 407, 1, 0, 0, 0, 109, 410, 1, 0, 0, 0, 111, 416, 1, 0, 0, 0,
		113, 427, 1, 0, 0, 0, 115, 116, 5, 105, 0, 0, 116, 117, 5, 109, 0, 0, 117,
		118, 5, 112, 0, 0, 118, 119, 5, 111, 0, 0, 119, 120, 5, 114, 0, 0, 120,
		121, 5, 116, 0, 0, 121, 2, 1, 0, 0, 0, 122, 123, 5, 112, 0, 0, 123, 124,
		5, 97, 0, 0, 124, 125, 5, 99, 0, 0, 125, 126, 5, 107, 0, 0, 126, 127, 5,
		97, 0, 0, 127, 128, 5, 103, 0, 0, 128, 129, 5, 101, 0, 0, 129, 4, 1, 0,
		0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 101, 0, 0, 132, 133, 5, 112,
		0, 0, 133, 134, 5, 101, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 116,
		0, 0, 136, 137, 5, 101, 0, 0, 137, 138, 5, 100, 0, 0, 138, 6, 1, 0, 0,
		0, 139, 140, 5, 109, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 112, 0,
		0, 142, 8, 1, 0, 0, 0, 143, 144, 5, 105, 0, 0, 144, 145, 5, 110, 0, 0,
		145, 146, 5, 116, 0, 0, 146, 147, 5, 51, 0, 0, 147, 148, 5, 50, 0, 0, 148,
		10, 1, 0, 0, 0, 149, 150, 5, 105, 0, 0, 150, 151, 5, 110, 0, 0, 151, 152,
		5, 116, 0, 0, 152, 153, 5, 54, 0, 0, 153, 154, 5, 52, 0, 0, 154, 12, 1,
		0, 0, 0, 155, 156, 5, 117, 0, 0, 156, 157, 5, 105, 0, 0, 157, 158, 5, 110,
		0, 0, 158, 159, 5, 116, 0, 0, 159, 160, 5, 51, 0, 0, 160, 161, 5, 50, 0,
		0, 161, 14, 1, 0, 0, 0, 162, 163, 5, 117, 0, 0, 163, 164, 5, 105, 0, 0,
		164, 165, 5, 110, 0, 0, 165, 166, 5, 116, 0, 0, 166, 167, 5, 54, 0, 0,
		167, 168, 5, 52, 0, 0, 168, 16, 1, 0, 0, 0, 169, 170, 5, 115, 0, 0, 170,
		171, 5, 105, 0, 0, 171, 172, 5, 110, 0, 0, 172, 173, 5, 116, 0, 0, 173,
		174, 5, 51, 0, 0, 174, 175, 5, 50, 0, 0, 175, 18, 1, 0, 0, 0, 176, 177,
		5, 115, 0, 0, 177, 178, 5, 105, 0, 0, 178, 179, 5, 110, 0, 0, 179, 180,
		5, 116, 0, 0, 180, 181, 5, 54, 0, 0, 181, 182, 5, 52, 0, 0, 182, 20, 1,
		0, 0, 0, 183, 184, 5, 102, 0, 0, 184, 185, 5, 105, 0, 0, 185, 186, 5, 120,
		0, 0, 186, 187, 5, 101, 0, 0, 187, 188, 5, 100, 0, 0, 188, 189, 5, 51,
		0, 0, 189, 190, 5, 50, 0, 0, 190, 22, 1, 0, 0, 0, 191, 192, 5, 102, 0,
		0, 192, 193, 5, 105, 0, 0, 193, 194, 5, 120, 0, 0, 194, 195, 5, 101, 0,
		0, 195, 196, 5, 100, 0, 0, 196, 197, 5, 54, 0, 0, 197, 198, 5, 52, 0, 0,
		198, 24, 1, 0, 0, 0, 199, 200, 5, 115, 0, 0, 200, 201, 5, 102, 0, 0, 201,
		202, 5, 105, 0, 0, 202, 203, 5, 120, 0, 0, 203, 204, 5, 101, 0, 0, 204,
		205, 5, 100, 0, 0, 205, 206, 5, 51, 0, 0, 206, 207, 5, 50, 0, 0, 207, 26,
		1, 0, 0, 0, 208, 209, 5, 115, 0, 0, 209, 210, 5, 102, 0, 0, 210, 211, 5,
		105, 0, 0, 211, 212, 5, 120, 0, 0, 212, 213, 5, 101, 0, 0, 213, 214, 5,
		100, 0, 0, 214, 215, 5, 54, 0, 0, 215, 216, 5, 52, 0, 0, 216, 28, 1, 0,
		0, 0, 217, 218, 5, 98, 0, 0, 218, 219, 5, 111, 0, 0, 219, 220, 5, 111,
		0, 0, 220, 221, 5, 108, 0, 0, 221, 30, 1, 0, 0, 0, 222, 223, 5, 115, 0,
		0, 223, 224, 5, 116, 0, 0, 224, 225, 5, 114, 0, 0, 225, 226, 5, 105, 0,
		0, 226, 227, 5, 110, 0, 0, 227, 228, 5, 103, 0, 0, 228, 32, 1, 0, 0, 0,
		229, 230, 5, 100, 0, 0, 230, 231, 5, 111, 0, 0, 231, 232, 5, 117, 0, 0,
		232, 233, 5, 98, 0, 0, 233, 234, 5, 108, 0, 0, 234, 235, 5, 101, 0, 0,
		235, 34, 1, 0, 0, 0, 236, 237, 5, 102, 0, 0, 237, 238, 5, 108, 0, 0, 238,
		239, 5, 111, 0, 0, 239, 240, 5, 97, 0, 0, 240, 241, 5, 116, 0, 0, 241,
		36, 1, 0, 0, 0, 242, 243, 5, 98, 0, 0, 243, 244, 5, 121, 0, 0, 244, 245,
		5, 116, 0, 0, 245, 246, 5, 101, 0, 0, 246, 247, 5, 115, 0, 0, 247, 38,
		1, 0, 0, 0, 248, 249, 5, 84, 0, 0, 249, 250, 5, 105, 0, 0, 250, 251, 5,
		109, 0, 0, 251, 252, 5, 101, 0, 0, 252, 253, 5, 115, 0, 0, 253, 254, 5,
		116, 0, 0, 254, 255, 5, 97, 0, 0, 255, 256, 5, 109, 0, 0, 256, 257, 5,
		112, 0, 0, 257, 40, 1, 0, 0, 0, 258, 259, 5, 68, 0, 0, 259, 260, 5, 117,
		0, 0, 260, 261, 5, 114, 0, 0, 261, 262, 5, 97, 0, 0, 262, 263, 5, 116,
		0, 0, 263, 264, 5, 105, 0, 0, 264, 265, 5, 111, 0, 0, 265, 266, 5, 110,
		0, 0, 266, 42, 1, 0, 0, 0, 267, 268, 5, 101, 0, 0, 268, 269, 5, 110, 0,
		0, 269, 270, 5, 117, 0, 0, 270, 271, 5, 109, 0, 0, 271, 44, 1, 0, 0, 0,
		272, 273, 5, 101, 0, 0, 273, 274, 5, 110, 0, 0, 274, 275, 5, 116, 0, 0,
		275, 276, 5, 105, 0, 0, 276, 277, 5, 116, 0, 0, 277, 278, 5, 121, 0, 0,
		278, 46, 1, 0, 0, 0, 279, 280, 5, 99, 0, 0, 280, 281, 5, 111, 0, 0, 281,
		282, 5, 109, 0, 0, 282, 283, 5, 112, 0, 0, 283, 284, 5, 111, 0, 0, 284,
		285, 5, 110, 0, 0, 285, 286, 5, 101, 0, 0, 286, 287, 5, 110, 0, 0, 287,
		288, 5, 116, 0, 0, 288, 48, 1, 0, 0, 0, 289, 290, 5, 59, 0, 0, 290, 50,
		1, 0, 0, 0, 291, 292, 5, 61, 0, 0, 292, 52, 1, 0, 0, 0, 293, 294, 5, 40,
		0, 0, 294, 54, 1, 0, 0, 0, 295, 296, 5, 41, 0, 0, 296, 56, 1, 0, 0, 0,
		297, 298, 5, 91, 0, 0, 298, 58, 1, 0, 0, 0, 299, 300, 5, 93, 0, 0, 300,
		60, 1, 0, 0, 0, 301, 302, 5, 123, 0, 0, 302, 62, 1, 0, 0, 0, 303, 304,
		5, 125, 0, 0, 304, 64, 1, 0, 0, 0, 305, 306, 5, 60, 0, 0, 306, 66, 1, 0,
		0, 0, 307, 308, 5, 62, 0, 0, 308, 68, 1, 0, 0, 0, 309, 310, 5, 46, 0, 0,
		310, 70, 1, 0, 0, 0, 311, 312, 5, 44, 0, 0, 312, 72, 1, 0, 0, 0, 313, 314,
		5, 58, 0, 0, 314, 74, 1, 0, 0, 0, 315, 316, 5, 43, 0, 0, 316, 76, 1, 0,
		0, 0, 317, 318, 5, 45, 0, 0, 318, 78, 1, 0, 0, 0, 319, 323, 5, 39, 0, 0,
		320, 322, 3, 81, 40, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323,
		324, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 326, 1, 0, 0, 0, 325, 323,
		1, 0, 0, 0, 326, 336, 5, 39, 0, 0, 327, 331, 5, 34, 0, 0, 328, 330, 3,
		81, 40, 0, 329, 328, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 332, 1, 0,
		0, 0, 331, 329, 1, 0, 0, 0, 332, 334, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0,
		334, 336, 5, 34, 0, 0, 335, 319, 1, 0, 0, 0, 335, 327, 1, 0, 0, 0, 336,
		80, 1, 0, 0, 0, 337, 342, 3, 83, 41, 0, 338, 342, 3, 85, 42, 0, 339, 342,
		3, 87, 43, 0, 340, 342, 8, 0, 0, 0, 341, 337, 1, 0, 0, 0, 341, 338, 1,
		0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 340, 1, 0, 0, 0, 342, 82, 1, 0, 0,
		0, 343, 344, 5, 92, 0, 0, 344, 345, 7, 1, 0, 0, 345, 346, 3, 107, 53, 0,
		346, 347, 3, 107, 53, 0, 347, 84, 1, 0, 0, 0, 348, 349, 5, 92, 0, 0, 349,
		350, 3, 105, 52, 0, 350, 351, 3, 105, 52, 0, 351, 352, 3, 105, 52, 0, 352,
		86, 1, 0, 0, 0, 353, 354, 5, 92, 0, 0, 354, 355, 7, 2, 0, 0, 355, 88, 1,
		0, 0, 0, 356, 357, 5, 116, 0, 0, 357, 358, 5, 114, 0, 0, 358, 359, 5, 117,
		0, 0, 359, 366, 5, 101, 0, 0, 360, 361, 5, 102, 0, 0, 361, 362, 5, 97,
		0, 0, 362, 363, 5, 108, 0, 0, 363, 364, 5, 115, 0, 0, 364, 366, 5, 101,
		0, 0, 365, 356, 1, 0, 0, 0, 365, 360, 1, 0, 0, 0, 366, 90, 1, 0, 0, 0,
		367, 371, 3, 93, 46, 0, 368, 371, 3, 95, 47, 0, 369, 371, 3, 97, 48, 0,
		370, 367, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 369, 1, 0, 0, 0, 371,
		92, 1, 0, 0, 0, 372, 376, 7, 3, 0, 0, 373, 375, 3, 103, 51, 0, 374, 373,
		1, 0, 0, 0, 375, 378, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0,
		0, 0, 377, 94, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 379, 383, 5, 48, 0, 0,
		380, 382, 3, 105, 52, 0, 381, 380, 1, 0, 0, 0, 382, 385, 1, 0, 0, 0, 383,
		381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 96, 1, 0, 0, 0, 385, 383, 1,
		0, 0, 0, 386, 387, 5, 48, 0, 0, 387, 389, 7, 1, 0, 0, 388, 390, 3, 107,
		53, 0, 389, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0,
		391, 392, 1, 0, 0, 0, 392, 98, 1, 0, 0, 0, 393, 398, 3, 101, 50, 0, 394,
		397, 3, 101, 50, 0, 395, 397, 3, 103, 51, 0, 396, 394, 1, 0, 0, 0, 396,
		395, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398, 399,
		1, 0, 0, 0, 399, 100, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 401, 402, 7, 4,
		0, 0, 402, 102, 1, 0, 0, 0, 403, 404, 7, 5, 0, 0, 404, 104, 1, 0, 0, 0,
		405, 406, 7, 6, 0, 0, 406, 106, 1, 0, 0, 0, 407, 408, 7, 7, 0, 0, 408,
		108, 1, 0, 0, 0, 409, 411, 7, 8, 0, 0, 410, 409, 1, 0, 0, 0, 411, 412,
		1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 414, 1, 0,
		0, 0, 414, 415, 6, 54, 0, 0, 415, 110, 1, 0, 0, 0, 416, 417, 5, 47, 0,
		0, 417, 418, 5, 47, 0, 0, 418, 422, 1, 0, 0, 0, 419, 421, 8, 9, 0, 0, 420,
		419, 1, 0, 0, 0, 421, 424, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 422, 423,
		1, 0, 0, 0, 423, 425, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 425, 426, 6, 55,
		1, 0, 426, 112, 1, 0, 0, 0, 427, 428, 5, 47, 0, 0, 428, 429, 5, 42, 0,
		0, 429, 433, 1, 0, 0, 0, 430, 432, 9, 0, 0, 0, 431, 430, 1, 0, 0, 0, 432,
		435, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 433, 431, 1, 0, 0, 0, 434, 436,
		1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 436, 437, 5, 42, 0, 0, 437, 438, 5, 47,
		0, 0, 438, 439, 1, 0, 0, 0, 439, 440, 6, 56, 1, 0, 440, 114, 1, 0, 0, 0,
		15, 0, 323, 331, 335, 341, 365, 370, 376, 383, 391, 396, 398, 412, 422,
		433, 2, 6, 0, 0, 0, 1, 0,
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

// kdsLexerInit initializes any static state used to implement kdsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewkdsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func KdsLexerInit() {
	staticData := &KdsLexerLexerStaticData
	staticData.once.Do(kdslexerLexerInit)
}

// NewkdsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewkdsLexer(input antlr.CharStream) *kdsLexer {
	KdsLexerInit()
	l := new(kdsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &KdsLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "kds.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// kdsLexer tokens.
const (
	kdsLexerIMPORT       = 1
	kdsLexerPACKAGE      = 2
	kdsLexerREPEATED     = 3
	kdsLexerMAP          = 4
	kdsLexerINT32        = 5
	kdsLexerINT64        = 6
	kdsLexerUINT32       = 7
	kdsLexerUINT64       = 8
	kdsLexerSINT32       = 9
	kdsLexerSINT64       = 10
	kdsLexerFIXED32      = 11
	kdsLexerFIXED64      = 12
	kdsLexerSFIXED32     = 13
	kdsLexerSFIXED64     = 14
	kdsLexerBOOL         = 15
	kdsLexerSTRING       = 16
	kdsLexerDOUBLE       = 17
	kdsLexerFLOAT        = 18
	kdsLexerBYTES        = 19
	kdsLexerTIMESTAMP    = 20
	kdsLexerDURATION     = 21
	kdsLexerENUM         = 22
	kdsLexerENTITY       = 23
	kdsLexerCOMPONENT    = 24
	kdsLexerSEMI         = 25
	kdsLexerEQ           = 26
	kdsLexerLP           = 27
	kdsLexerRP           = 28
	kdsLexerLB           = 29
	kdsLexerRB           = 30
	kdsLexerLC           = 31
	kdsLexerRC           = 32
	kdsLexerLT           = 33
	kdsLexerGT           = 34
	kdsLexerDOT          = 35
	kdsLexerCOMMA        = 36
	kdsLexerCOLON        = 37
	kdsLexerPLUS         = 38
	kdsLexerMINUS        = 39
	kdsLexerSTR_LIT      = 40
	kdsLexerBOOL_LIT     = 41
	kdsLexerINT_LIT      = 42
	kdsLexerIDENTIFIER   = 43
	kdsLexerWS           = 44
	kdsLexerLINE_COMMENT = 45
	kdsLexerCOMMENT      = 46
)
