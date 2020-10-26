package errors

type ErrorType int

const (
    NotNogE ErrorType = iota

)

type nogError struct{
  Type ErrorType
  msg string
}

func NewNotNog(msg string) (*nogError) {
  e := new(nogError)
  e.msg=msg
  e.Type=NotNogE

  return e

}

func (m *nogError) Error() string {
    return m.msg
}

func NotNog(e interface{}) (bool) {

  if e==nil {return false}
  v, ok := e.(*nogError)
  if !ok { return false}

  return v.Type==NotNogE

}
