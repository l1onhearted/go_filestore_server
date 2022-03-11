package meta

type FileMeta struct { //文件元信息结构
	FileSha1 string
	FileName string
	FileSize int64  //文件大小
	Location string //存储位置
	UploadAt string //上传时间
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//新增更新文件元信息
func UpdateFileMeta(fmeta FileMeta) { 
	fileMetas[fmeta.FileSha1] = fmeta
}

//通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta { 
	return fileMetas[fileSha1]
}
