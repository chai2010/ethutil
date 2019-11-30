// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import "testing"

import (
	"encoding/json"
	"io/ioutil"
)

func TestKeyStoreDecryptKey(t *testing.T) {
	for _, tm := range []map[string]interface{}{
		tLoadTestdata(t, "v1_test_vector.json"),
		tLoadTestdata(t, "v3_test_vector.json"),
	} {
		for tname, m := range tm {
			m := m.(map[string]interface{})
			mJson := m["json"].(map[string]interface{})

			mJsonData, err := json.Marshal(mJson)
			if err != nil {
				t.Fatal(tname, err)
			}

			id, priv, err := KeyStoreDecryptKey(mJsonData, m["password"].(string))
			tAssert(t, err == nil, tname, err)
			tAssert(t, id == mJson["id"].(string), tname, id, priv)
			tAssert(t, m["priv"] == priv, tname, id, priv)
		}
	}
}

// 加载 testdata 目录下的文件
func tLoadTestdata(tb testing.TB, filename string) map[string]interface{} {
	data, err := ioutil.ReadFile("./testdata/" + filename)
	if err != nil {
		tb.Fatal(err)
	}

	var m = make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		tb.Fatal(err)
	}
	return m
}
