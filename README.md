# twitter_clone


-Sırasıyla aşağıdaki komutları çalıştırınız.
- `git clone https://github.com/onurcanari/twitter_clone.git` 
- `docker-compose up --build` 

-http://localhost:8080 de proje çalışmaya başlamış olacaktır.
CORS bazı isteklerde sorun çıkardığından siteye giriş yapabilmek için tarayıcıda CORS u devre dışı bırakmak gerekmektedir. Şu şekilde yapılabilir.

Windows
- `chrome.exe --user-data-dir="C:/Chrome dev session" --disable-web-security`

OSX
- `open -n -a /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --args --user-data-dir="/tmp/chrome_dev_test" --disable-web-security`

Linux
- `google-chrome --disable-web-security`
