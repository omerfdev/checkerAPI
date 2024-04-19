# Karakter Karşılaştırma API

Bu basit API, Rick and Morty API'sinden iki farklı karakterin verilerini çeker ve bu karakterleri karşılaştırır. Eğer iki karakterin verileri aynı ise, HTTP status kodu 200 (OK) döner, aksi halde 401 (Unauthorized) döner.

## Kurulum

1. Bu kodu çalıştırmak için öncelikle Go programlama dilinin yüklü olması gerekmektedir.
2. Kodu çalıştırmak için terminalde `go run main.go` komutunu çalıştırın.
3. Tarayıcınızdan veya bir API test aracından `http://localhost:8082/check_endpoint` adresine istek gönderin.

## Kullanım

- API, `/check_endpoint` endpoint'i üzerinden iki karakterin verilerini karşılaştırır.
- Her istek, iki farklı karakterin verilerini alır ve karşılaştırır.
- Eğer karakter verileri aynı ise, HTTP status kodu 200 (OK) döner, aksi halde 401 (Unauthorized) döner.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
