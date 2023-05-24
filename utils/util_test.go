package utils

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestExt(t *testing.T) {
	ext := Ext("sqlc.yaml")
	t.Logf("ext: %s\n", ext)
}

func TestMime(t *testing.T) {
	var exts = []string{
		"image/avif",
		"text/css; charset=utf-8",
		"image/jpeg",
		"application/pdf",
		"text/xml; charset=utf-8",
		"image/webp",
	}

	// var wg sync.WaitGroup
	for i := 0; i < len(exts); i++ {
		// wg.Add(1)
		go func(i int) {
			// defer wg.Done()
			ext, ok := GetInstance()[exts[i]]
			t.Logf("ext: %s,ok: %t", ext, ok)
		}(i)
		time.Sleep(1 * time.Second)
	}

	// wg.Wait()
}

func TestUrl(t *testing.T) {
	for _, v := range []string{
		"/usr/local/etc",
		"http://www.baidu.com",
		`https://ts1.cn.mm.bing.net/th/id/R-C.987f582c510be58755c4933cda68d525?rik=C0D21hJDYvXosw&riu=http%3a%2f%2fimg.pconline.com.cn%2fimages%2fupload%2fupc%2ftx%2fwallpaper%2f1305%2f16%2fc4%2f20990657_1368686545122.jpg&ehk=netN2qzcCVS4ALUQfDOwxAwFcy41oxC%2b0xTFvOYy5ds%3d&risl=&pid=ImgRaw&r=0`,
	} {
		b := InValidUrl(v)
		t.Logf("b: %t", b)
	}
}

func TestGetUrl(t *testing.T) {
	var u = `https://ts1.cn.mm.bing.net/th/id/R-C.987f582c510be58755c4933cda68d525?rik=C0D21hJDYvXosw&riu=http%3a%2f%2fimg.pconline.com.cn%2fimages%2fupload%2fupc%2ftx%2fwallpaper%2f1305%2f16%2fc4%2f20990657_1368686545122.jpg&ehk=netN2qzcCVS4ALUQfDOwxAwFcy41oxC%2b0xTFvOYy5ds%3d&risl=&pid=ImgRaw&r=0`

	resp, err := http.Get(u)
	require.NoError(t, err)
	t.Logf("len: %f", float64(resp.ContentLength/(1024*1024)))
}
