package mimes

import (
	"errors"
)

var mimeTypes map[string]string
var revMimeTypes map[string]string //reversed

func init() {
	mimeTypes = make(map[string]string)
	revMimeTypes = make(map[string]string)

	mimeTypes[".aac"] = "audio/aac"
	mimeTypes[".abw"] = "application/x-abiword"
	mimeTypes[".arc"] = "application/octet-stream"
	mimeTypes[".avi"] = "video/x-msvideo"
	mimeTypes[".azw"] = "application/vnd.amazon.ebook"
	mimeTypes[".bin"] = "application/octet-stream"
	mimeTypes[".bz"] = "application/x-bzip"
	mimeTypes[".bz2"] = "application/x-bzip2"
	mimeTypes[".csh"] = "application/x-csh"
	mimeTypes[".css"] = "text/css"
	mimeTypes[".csv"] = "text/csv"
	mimeTypes[".doc"] = "application/msword"
	mimeTypes[".docx"] = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	mimeTypes[".eot"] = "application/vnd.ms-fontobject"
	mimeTypes[".epub"] = "application/epub+zip"
	mimeTypes[".es"] = "application/ecmascript"
	mimeTypes[".gif"] = "image/gif"
	mimeTypes[".htm"] = "text/html"
	mimeTypes[".html"] = "text/html"
	mimeTypes[".ico"] = "image/x-icon"
	mimeTypes[".ics"] = "text/calendar"
	mimeTypes[".jar"] = "application/java-archive"
	mimeTypes[".jpeg"] = "image/jpeg"
	mimeTypes[".jpg"] = "image/jpeg"
	mimeTypes[".js"] = "application/javascript"
	mimeTypes[".json"] = "application/json"
	mimeTypes[".mid"] = "audio/midi audio/x-midi"
	mimeTypes[".midi"] = "audio/midi audio/x-midi"
	mimeTypes[".mpeg"] = "video/mpeg"
	mimeTypes[".mpkg"] = "application/.apple.installer+xml"
	mimeTypes[".odp"] = "application/vnd.oasis.opendocument.presentation"
	mimeTypes[".ods"] = "application/vnd.oasis.opendocument.spreadsheet"
	mimeTypes[".odt"] = "application/vnd.oasis.opendocument.text"
	mimeTypes[".oga"] = "audio/ogg"
	mimeTypes[".ogv"] = "video/ogg"
	mimeTypes[".ogx"] = "application/ogg"
	mimeTypes[".otf"] = "font/otf"
	mimeTypes[".png"] = "image/png"
	mimeTypes[".pdf"] = "application/pdf"
	mimeTypes[".ppt"] = "application/vnd.ms-powerpoint"
	mimeTypes[".pptx"] = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	mimeTypes[".rar"] = "application/x-rar-compressed"
	mimeTypes[".rtf"] = "application/rtf"
	mimeTypes[".sh"] = "application/x-sh"
	mimeTypes[".svg"] = "image/svg+xml"
	mimeTypes[".swf"] = "application/x-shockwave-flash"
	mimeTypes[".tar"] = "application/x-tar"
	mimeTypes[".tif"] = "image/tiff"
	mimeTypes[".tiff"] = "image/tiff"
	mimeTypes[".ts"] = "application/typescript"
	mimeTypes[".ttf"] = "font/ttf"
	mimeTypes[".vsd"] = "application/vnd"
	mimeTypes[".visio"] = ".wav audio/wav"
	mimeTypes[".wasm"] = "application/wasm"
	mimeTypes[".weba"] = "audio/webm"
	mimeTypes[".webm"] = "video/webm"
	mimeTypes[".webp"] = "image/webp"
	mimeTypes[".woff"] = "font/woff"
	mimeTypes[".woff2"] = "font/woff2"
	mimeTypes[".xhtml"] = "application/xhtml+xml"
	mimeTypes[".xls"] = "application/vnd.ms-excel"
	mimeTypes[".xlsx"] = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	mimeTypes[".xml"] = "application/xml"
	mimeTypes[".xul"] = "application/vnd.mozilla.xul+xml"
	mimeTypes[".zip"] = "application/zip"
	mimeTypes[".3gp"] = "video/3gpp"
	mimeTypes[".3g2"] = "video/3gpp2"
	mimeTypes[".7z"] = "application/x-7z-compressed"

	for key, value := range mimeTypes {
		revMimeTypes[value] = key
	}
}

func Get(ext string) (mime string, err error) {

	mime, exists := mimeTypes[ext]

	if !exists {
		mime = "text/plain"
		err = errors.New("extension not found")
	}

	return
}

func GetExt(mime string) (ext string, err error) {

	ext, exists := revMimeTypes[mime]

	if !exists {
		err = errors.New("mime type not found")
	}

	return
}
