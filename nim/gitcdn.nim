import jester, httpclient, re, ospaths, mimetypes, strutils

routes:
  get re"(.*)":
    cond: request.matches[0].len > 1
    try:
      resp getContent("https://raw.githubusercontent.com" & request.matches[0]),
           getMimeType(newMimeTypes(), strip(splitFile(request.matches[0]).ext, true, false, {'.'}))
    except:
      resp Http404