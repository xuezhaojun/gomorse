package morse

var dict map[byte]string = map[byte]string{
	'A':  ".-",
	'B':  "-...",
	'C':  "-.-.",
	'D':  "-..",
	'E':  ".",
	'F':  "..-.",
	'G':  "--.",
	'H':  "....",
	'I':  "..",
	'J':  ".---",
	'K':  "-.-",
	'L':  ".-..",
	'M':  "--",
	'N':  "-.",
	'O':  "---",
	'P':  ".--.",
	'Q':  "--.-",
	'R':  ".-.",
	'S':  "...",
	'T':  "-",
	'U':  "..-",
	'V':  "...-",
	'W':  ".--",
	'X':  "-..-",
	'Y':  "-.--",
	'Z':  "--..",
	'1':  ".----",
	'2':  "..---",
	'3':  "...--",
	'4':  "....-",
	'5':  ".....",
	'6':  "-....",
	'7':  "--...",
	'8':  "---..",
	'9':  "----.",
	'0':  "-----",
	'.':  ".-.-.-",
	':':  "---...",
	',':  "--..--",
	';':  "-.-.-.",
	'?':  "..--..",
	'=':  "-...-",
	'\'': ".----.",
	'/':  "-..-.",
	'!':  "-.-.--",
	'-':  "-....-",
	'_':  "..--.-",
	'"':  ".-..-.",
	'(':  "-.--.",
	')':  "-.--.-",
	'$':  "...-..-",
	'&':  ".-...",
	'@':  ".--.-.",
	'+':  ".-.-.",
	' ':  " ",
}

var redict map[string]byte = map[string]byte{
	"--.":     'G',
	".....":   '5',
	"-....":   '6',
	"--":      'M',
	"---":     'O',
	".--.":    'P',
	"-":       'T',
	".----":   '1',
	"-..":     'D',
	"-.":      'N',
	"...-..-": '$',
	"...":     'S',
	"..--.-":  '_',
	".-":      'A',
	"----.":   '9',
	"-...-":   '=',
	".-...":   '&',
	"-----":   '0',
	"-..-.":   '/',
	"..":      'I',
	"-.--":    'Y',
	".-..-.":  '"',
	"-...":    'B',
	".-.":     'R',
	"--...":   '7',
	"---..":   '8',
	"-.-.--":  '!',
	".--.-.":  '@',
	"..-.":    'F',
	"-..-":    'X',
	"--..":    'Z',
	".-.-.-":  '.',
	"-.--.-":  ')',
	"-.-.":    'C',
	".":       'E',
	".--":     'W',
	"...--":   '3',
	"-.--.":   '(',
	".----.":  '\'',
	".---":    'J',
	"-.-":     'K',
	"--.-":    'Q',
	"..---":   '2',
	"--..--":  ',',
	"....":    'H',
	"....-":   '4',
	"-.-.-.":  ';',
	"..--..":  '?',
	"...-":    'V',
	"---...":  ':',
	".-.-.":   '+',
	".-..":    'L',
	"..-":     'U',
	"-....-":  '-',
	" ":       ' ',
}
