package jsontool

import (
	"fmt"
	"reflect"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type jsontool struct {
	json     jsoniter.API
	typelist reflect.Type
	typemap  reflect.Type
	Params   map[string]interface{}
}

func JsonToolInstance(msg interface{}) (*jsontool, error) {
	var err error
	intsance := &jsontool{json: jsoniter.ConfigCompatibleWithStandardLibrary,
		typelist: reflect.TypeOf([]interface{}{1}),
		typemap:  reflect.TypeOf(map[string]interface{}{"1": 1}),
		Params:   make(map[string]interface{})}
	switch msg.(type) {
	case *strings.Reader:
		//fmt.Printf("is *strings.Reader\n")
		decoder := intsance.json.NewDecoder(msg.(*strings.Reader))
		err = decoder.Decode(&(intsance.Params))
	case []byte:
		//fmt.Printf("is []byte\n")
		err = intsance.json.Unmarshal(msg.([]byte), &(intsance.Params))
	case string:
		//fmt.Printf("is string\n")
		err = intsance.json.UnmarshalFromString(msg.(string), &(intsance.Params))
	}
	return intsance, err
}

func (jtool *jsontool) GetValue(keys ...interface{}) interface{} {
	var value interface{}
	value = jtool.Params

	for _, key := range keys {
		switch key.(type) {
		case int:
			if reflect.TypeOf(value) != jtool.typelist {
				return nil
			}
			value = value.([]interface{})
			if len(value.([]interface{})) > key.(int) {
				value = value.([]interface{})[key.(int)]
			} else {
				return nil
			}
			//fmt.Printf("int=%v,%v\n", key, value)
		case string:
			if reflect.TypeOf(value) != jtool.typemap {
				return nil
			}
			value = value.(map[string]interface{})[key.(string)]
			//fmt.Printf("string=%v,%v\n", key, value)
		default:
			fmt.Printf("不支持的数据类型\n")
		}
	}
	return value
}

func test() {
	//msg := strings.NewReader(`{"branch":[{"name":"beta"}],"change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`)
	//msg := strings.NewReader(`error test "branch":[{"name":"beta"}],"change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`)
	//msg := []byte(`{"branch":[{"name":"beta"}],"change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`)
	msg := `{"branch":[{"name":"beta"}],"change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`
	jtool, err := JsonToolInstance(msg)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("111==>>%s,%s,%s\n", jtool.Params["branch"].([]interface{})[0].(map[string]interface{})["name"], jtool.Params["channel"], jtool.Params["md5"])
		value := jtool.GetValue("branch", 0, "name")
		fmt.Printf("value=%v\n", value)
		//fmt.Printf("branch type=%t,%t\n", jtool.typelist == jtool.typemap, jtool.typelist.Kind() == reflect.Array)
		value = jtool.GetValue("0", 0, "name")
		fmt.Printf("value=%v\n", value)
		value = jtool.GetValue("branch", 0, "test")
		fmt.Printf("value=%v\n", value)

	}
}
