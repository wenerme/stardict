package codec

import "time"

const InfoMagic = "StarDict's dict ifo file"

type Dict struct {
	Info    *DictInfo
	Entries []*DictEntry
}

type DictInfo struct {
	Header           string // magic StarDict's dict ifo file
	Version          string // "2.4.2" or "3.0.0"
	WordCount        int
	SynWordCount     int
	IndexFileSize    int
	IndexOffsetBits  int // since 3.0.0 Bits pre offset, 32/64
	BookName         string
	Description      string // <br> for new line
	Date             time.Time
	SameTypeSequence string
	DictType         string

	Author  string
	Email   string
	Website string
}

type DictIndex struct {
	Entries []*DictEntry
}
type DictEntry struct {
	Word string // a utf-8 string terminated by '\0'.
	// word data's offset in .dict file<br>
	// If the version is "3.0.0" and "idxoffsetbits=64", word_data_offset will be 64-bits unsigned number in network byte order.
	// Otherwise it will be 32-bits.
	Offset uint64
	// word data's total size in .dict file<br>
	// word_data_size should be 32-bits unsigned number in network byte order.
	Size     int
	Synonyms []string

	Contents []*EntryContent
}

/*
{1}. Files
Every dictionary consists of these files:
(1). somedict.ifo
(2). somedict.idx or somedict.idx.gz
(3). somedict.dict or somedict.dict.dz
(4). somedict.syn (optional)

StarDict search for dictionaries in the following predefined directories:
1) gStarDictDataDir + "/dic",
2) "/usr/share/stardict/dic",
3) g_get_home_dir() + "/.stardict/dic".
*/

// Lower-case characters signify that a field's size is determined by a
// terminating '\0', while upper-case characters indicate that the data
// begins with a network byte-ordered guint32 that gives the length of
// the following data's size (NOT the whole size which is 4 bytes bigger).
type ContentType rune

const (
	// Word's pure text meaning.
	// The data should be a utf-8 string ending with '\0'.
	NULL_TERMINAL_TEXT ContentType = 'm'

	// Word's pure text meaning.
	// The data is NOT a utf-8 string, but is instead a string in locale
	// encoding, ending with '\0'. Sometimes using this type will save disk
	// space, but its use is discouraged. This is only a idea.
	LOCALE_TEXT ContentType = 'l'

	// A utf-8 string which is marked up with the Pango text markup language.
	// For more information about this markup language, See the "Pango
	// Reference Manual."
	// You might have it installed locally at:
	// file:///usr/share/gtk-doc/html/pango/PangoMarkupFormat.html
	// Online:
	// http://library.gnome.org/devel/pango/stable/PangoMarkupFormat.html
	PANGO_TEXT ContentType = 'g'

	// English phonetic string.
	// The data should be a utf-8 string ending with '\0'.
	//
	// Here are some utf-8 phonetic characters:
	//
	// Î¸ÊƒÅ‹Ê§Ã°Ê’Ã¦Ä±ÊŒÊŠÉ’É›É™É‘ÉœÉ”ËŒËˆËË‘á¹ƒá¹‡á¸·
	// Ã¦É‘É’ÊŒÓ™Ñ”Å‹vÎ¸Ã°ÊƒÊ’ÉšËÉ¡ËËŠË‹
	ENGLISH_PHONETIC ContentType = 't'

	// A utf-8 string which is marked up with the xdxf language.
	// StarDict have these extension:
	// <rref> can have "type" attribute, it can be "image", "sound", "video"
	// and "attach".
	// <kref> can have "k" attribute.
	//
	// http://xdxf.sourceforge.net
	XDXF_MARKUP ContentType = 'x'
	// Chinese YinBiao or Japanese KANA.
	// The data should be a utf-8 string ending with '\0'.
	YINBIAO ContentType = 'y'
	// KingSoft PowerWord's data. The data is a utf-8 string ending with '\0'.
	// It is in XML format.
	KINGSOFT_XML ContentType = 'k'
	// MediaWiki markup language.
	// http://meta.wikimedia.org/wiki/Help:Editing#The_wiki_markup
	MEDIAWIKI_MARKUP ContentType = 'w'
	// Html codes.
	HTML ContentType = 'h'
	// WordNet data.
	WORDNET ContentType = 'n'
	// Resource file list.
	// The content can be:
	// img:pic/example.jpg	// Image file
	// snd:apple.wav		// Sound file
	// vdo:film.avi		// Video file
	// att:file.bin		// Attachment file
	// More than one line is supported as a list of available files.
	// StarDict will find the files in the Resource Storage.
	// The image will be shown, the sound file will have a play button.
	// You can "save as" the attachment file and so on.
	// The file list must be a utf-8 string ending with '\0'.
	// Use '\n' for separating new lines.
	// Use '/' character as directory separator.
	RESOURCE ContentType = 'r'
	// wav file.
	// The data begins with a network byte-ordered guint32 to identify the wav
	// file's size, immediately followed by the file's content.
	// This is only a idea, it is better to use 'r' Resource file list in most
	// case.
	WAV_FILE ContentType = 'W'
	// Picture file.
	// The data begins with a network byte-ordered guint32 to identify the picture
	// file's size, immediately followed by the file's content.
	// This feature is implemented, as stardict-advertisement-plugin needs it.
	// Anyway, it is better to use 'r' Resource file list in most case.
	PICTURE_FILE ContentType = 'P'
	// this type identifier is reserved for experimental extensions.
	RESERVED ContentType = 'X'
)
