package api

import "net/http"

type Jenkins struct {
	Host string
	Password string
	Client http.Client
	User string
	Token string
}

func NewJenkins(user,password,host,token string) *Jenkins{
	return &Jenkins{
		Host:     host,
		Password: password,
		Client:   http.Client{},
		User:     user,
		Token:    token,
	}
}


//func (r *Requester) Do(ar *APIRequest, responseStruct interface{}, options ...interface{}) (*http.Response, error) {
//	if !strings.HasSuffix(ar.Endpoint, "/") && ar.Method != "POST" {
//		ar.Endpoint += "/"
//	}
//
//	fileUpload := false
//	var files []string
//	URL, err := url.Parse(r.Base + ar.Endpoint + ar.Suffix)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, o := range options {
//		switch v := o.(type) {
//		case map[string]string:
//
//			querystring := make(url.Values)
//			for key, val := range v {
//				querystring.Set(key, val)
//			}
//
//			URL.RawQuery = querystring.Encode()
//		case []string:
//			fileUpload = true
//			files = v
//		}
//	}
//	var req *http.Request
//
//	if fileUpload {
//		body := &bytes.Buffer{}
//		writer := multipart.NewWriter(body)
//		for _, file := range files {
//			fileData, err := os.Open(file)
//			if err != nil {
//				Error.Println(err.Error())
//				return nil, err
//			}
//
//			part, err := writer.CreateFormFile("file", filepath.Base(file))
//			if err != nil {
//				Error.Println(err.Error())
//				return nil, err
//			}
//			if _, err = io.Copy(part, fileData); err != nil {
//				return nil, err
//			}
//			defer fileData.Close()
//		}
//		var params map[string]string
//		json.NewDecoder(ar.Payload).Decode(&params)
//		for key, val := range params {
//			if err = writer.WriteField(key, val); err != nil {
//				return nil, err
//			}
//		}
//		if err = writer.Close(); err != nil {
//			return nil, err
//		}
//		req, err = http.NewRequest(ar.Method, URL.String(), body)
//		if err != nil {
//			return nil, err
//		}
//		req.Header.Set("Content-Type", writer.FormDataContentType())
//	} else {
//		req, err = http.NewRequest(ar.Method, URL.String(), ar.Payload)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	if r.BasicAuth != nil {
//		req.SetBasicAuth(r.BasicAuth.Username, r.BasicAuth.Password)
//	}
//
//	for k := range ar.Headers {
//		req.Header.Add(k, ar.Headers.Get(k))
//	}
//	fmt.Printf("%#v\n",*req.URL)
//	if response, err := r.Client.Do(req); err != nil {
//		return nil, err
//	} else {
//		errorText := response.Header.Get("X-Error")
//		if errorText != "" {
//			return nil, errors.New(errorText)
//		}
//		switch responseStruct.(type) {
//		case *string:
//			return r.ReadRawResponse(response, responseStruct)
//		default:
//			return r.ReadJSONResponse(response, responseStruct)
//		}
//
//	}
//
//}