type File <br/>
   &nbsp;&nbsp;func ReadFile(path string) *File <br/>
   &nbsp;&nbsp;func Data(f FileInter) []byte <br/>
   &nbsp;&nbsp;func (d *File) DataBuf() []byte <br/>


func WriteDataFile(data []byte, nameFile string) (string, error) <br/>

type FileInter interface {
   &nbsp;&nbsp;&nbsp;&nbsp;DataBuf() []byte
} <br/>