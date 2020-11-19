type File
   &nbsp;&nbsp;func ReadFile(path string) *File
   &nbsp;&nbsp;func Data(f FileInter) []byte
   &nbsp;&nbsp;func (d *File) DataBuf() []byte


func WriteDataFile(data []byte, nameFile string) (string, error)

type FileInter interface {
   &nbsp;&nbsp;&nbsp;&nbsp;DataBuf() []byte
}