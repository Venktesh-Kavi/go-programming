Quick Note:

* In IO if func ReadFull(r Reader, b []byte) (int, error), has a byte buffer length greater than the
  src buffer len. Then the full buffer is not utilised resulting in unexpected EOF exceptions as per
  docs. The use of ReadFull is to use the full buffer provided by the user during readFull call. Consider ReadAll for dynamic use cases;