# Rise
String metric based image crawler

[![Build Status](https://travis-ci.org/if1live/rise.svg?branch=master)](https://travis-ci.org/if1live/rise)

## Tutorial

I want to download some images from http://bcy.net/coser/detail/47840/602658.

```bash
./rise -cmd=show -uri=http://bcy.net/coser/detail/47840/602658

cid=0 : /Public/Image/zhipin/tip-zhipin.png
cid=1 : /Public/Image/sub-nav/logo-cos-channel.png
cid=2 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=2 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=3 : /Public/Image/payment/giftDetail.png
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg/w650
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/215c6f5032ee11e6a43b0f9416af7213.jpg/w650
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/2bfe252032ee11e6a43b0f9416af7213.jpg/w650
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/0d79a52032ee11e6a43b0f9416af7213.jpg/w650
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/1d5edbe032ee11e6a43b0f9416af7213.jpg/w650
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/12ae801032ee11e6a43b0f9416af7213.jpg/w650
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/17f8b7c032ee11e6a43b0f9416af7213.jpg/w650
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/26e97c6032ee11e6a43b0f9416af7213.jpg/w650
cid=5 : http://img5.bcyimg.com/coser/47840/post/17815/a3f10c30348f11e6b89e1967e57cfb1f.jpg/w650
cid=5 : http://img5.bcyimg.com/coser/47840/post/17815/ac8f9f00348f11e6b89e1967e57cfb1f.jpg/w650
cid=5 : http://img9.bcyimg.com/coser/47840/post/17815/7a7ae1a0348f11e6b89e1967e57cfb1f.jpg/w650
cid=5 : http://img9.bcyimg.com/coser/47840/post/17815/852ee6f0348f11e6b89e1967e57cfb1f.jpg/w650
cid=5 : http://img5.bcyimg.com/coser/47840/post/17815/8c4e2cc0348f11e6b89e1967e57cfb1f.jpg/w650
cid=5 : http://img9.bcyimg.com/coser/47840/post/17815/939fa620348f11e6b89e1967e57cfb1f.jpg/w650
cid=5 : http://img9.bcyimg.com/coser/47840/post/17815/98e4fbd0348f11e6b89e1967e57cfb1f.jpg/w650
cid=6 : http://img5.bcyimg.com/coser/47840/post/17815/49563420349011e6aa3b1bfee02be998.jpg/w650
cid=7 : http://img5.bcyimg.com/coser/47840/post/17813/1a60f67032f411e6941727176b80f3f3.jpg/w650
cid=8 : http://img9.bcyimg.com/coser/47840/post/17813/e11d62e032f311e6941727176b80f3f3.jpg/w650
cid=8 : http://img9.bcyimg.com/coser/47840/post/17813/246d4a1032f411e6941727176b80f3f3.jpg/w650
cid=8 : http://img5.bcyimg.com/coser/47840/post/17813/ea0cfbe032f311e6941727176b80f3f3.jpg/w650
cid=9 : http://img5.bcyimg.com/coser/47840/post/17818/c2cdcf1036e411e6b71051ef0ffbe5d8.jpg/w650
cid=9 : http://img9.bcyimg.com/coser/47840/post/17818/d01aa7b036e411e6b71051ef0ffbe5d8.jpg/w650
cid=9 : http://img9.bcyimg.com/coser/47840/post/17818/e0c2923036e411e6b71051ef0ffbe5d8.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/f269992032f211e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/f8b9a81032f211e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/fd80a60032f211e6941727176b80f3f3.jpg/w650
cid=10 : http://img5.bcyimg.com/coser/47840/post/17813/0226876032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/0653412032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/0a88d48032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img5.bcyimg.com/coser/47840/post/17813/0e07e92032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/1270e89032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/1660242032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/1a31295032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img5.bcyimg.com/coser/47840/post/17813/1f32712032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img5.bcyimg.com/coser/47840/post/17813/2337097032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/2801ff0032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/2cc9722032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/30fbd13032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img5.bcyimg.com/coser/47840/post/17813/34446c8032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/390b918032f311e6941727176b80f3f3.jpg/w650
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/3dd4d96032f311e6941727176b80f3f3.jpg/w650
cid=11 : http://img9.bcyimg.com/coser/47840/post/17818/f0241a0036e411e6b71051ef0ffbe5d8.jpg/w650
cid=11 : http://img9.bcyimg.com/coser/47840/post/17818/f7b645e036e411e6b71051ef0ffbe5d8.jpg/w650
cid=11 : http://img9.bcyimg.com/coser/47840/post/17818/0591c9f036e511e6b71051ef0ffbe5d8.jpg/w650
cid=11 : http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg/w650
cid=12 : /Public/Image/app.png
cid=13 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=14 : http://img9.bcyimg.com/coser/47840/post/4blp/8a4be4a027cc11e6883c1736fee6edeb.jpg/2X2
cid=15 : http://img5.bcyimg.com/coser/47840/post/17814/1687212033c411e6b1c1e563932fa8f9.jpg/2X2
cid=16 : http://img9.bcyimg.com/coser/47840/post/4bm2/22e159a041ec11e6a48aa1b424cea3b5.jpg/2X3
cid=17 : http://img9.bcyimg.com/coser/47840/post/17818/28a8bba036eb11e6b88433af95948698.jpg/2X3
cid=18 : http://img5.bcyimg.com/coser/47840/post/17814/1687212033c411e6b1c1e563932fa8f9.jpg/2X3
cid=19 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg/2X3
cid=20 : http://img9.bcyimg.com/coser/47840/post/4blp/8a4be4a027cc11e6883c1736fee6edeb.jpg/2X3
cid=21 : http://img9.bcyimg.com/editor/flag/c04hy/0e964d00c80d11e690af7feef9a49a4a.jpg
cid=22 : http://img9.bcyimg.com/editor/flag/1789t/0f644a60862d11e6806365efe9e0d6f6.jpg
cid=23 : http://user.bcyimg.com/Public/Upload/avatar/1971889/8cb37ce0642c41609f0e98a3674aded8/middle.jpg
cid=24 : http://user.bcyimg.com/Public/Upload/avatar/1971495/eefef6831d414dd1aa02134788b01497/middle.jpg
cid=25 : /Public/Image/user_pic_middle.png
cid=26 : http://user.bcyimg.com/Public/Upload/avatar/829981/d345cba216ec43de96ec3a46967585b9/middle.jpg
cid=26 : http://user.bcyimg.com/Public/Upload/avatar/829981/d345cba216ec43de96ec3a46967585b9/middle.jpg
cid=27 : /Public/Image/user_pic_middle.png
cid=28 : http://user.bcyimg.com/Public/Upload/avatar/1722841/5771c24e17f4453392344f3b2d485273/middle.jpg
cid=29 : http://user.bcyimg.com/Public/Upload/avatar/1029245/42a6d27d3b0a4bb9aad95554d8bffabb/middle.jpg
cid=30 : /Public/Image/user_pic_middle.png
cid=31 : http://user.bcyimg.com/Public/Upload/avatar/1032988/f63d8ef8ac924dfd920a9c56f7c78563/middle.jpg
cid=32 : /Public/Image/user_pic_middle.png
cid=33 : http://user.bcyimg.com/Public/Upload/avatar/955015/29a05bd0218a4aa6a08a269c29a5daa7/middle.jpg
cid=34 : /Public/Image/user_pic_middle.png
cid=34 : /Public/Image/user_pic_middle.png
cid=35 : http://user.bcyimg.com/Public/Upload/avatar/1200649/feac93a618b84096a8d713ce65b86f78/middle.jpg
cid=36 : http://img9.bcyimg.com/article/post/c04hs/49192870c35911e685cda78b1017df95.jpg/w230
cid=36 : http://img9.bcyimg.com/article/post/c04hs/3d2724e0c35911e685cda78b1017df95.jpg/w230
cid=37 : http://img9.bcyimg.com/article/post/c04hq/d300f930c1d511e6a6ecab0bcfeee30f.jpg/w230
cid=38 : http://img9.bcyimg.com/article/post/c04hy/e7b73b50c7fc11e6b37e0f01456a9728.jpg/w230
cid=39 : http://img9.bcyimg.com/article/post/c04hx/755c5590c75211e69609e1e8c1b0018d.jpg/w230
cid=40 : /Public/Image/app.png
cid=41 : /Public/Image/qr-cltAndr.gif
```

I want download those links.

* http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg/w650 (6th link)
* .....
* http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg/w650

Those links has suffix, "/w650". 
"/w650" is a parameter to create thumbnail, width 650px.
I want url without parameter. use `-filter=1` to slice extra strings after extension.


```
./rise -cmd=show -uri=http://bcy.net/coser/detail/47840/602658 -filter=1

cid=0 : /Public/Image/zhipin/tip-zhipin.png
cid=1 : /Public/Image/sub-nav/logo-cos-channel.png
cid=2 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=2 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=3 : /Public/Image/payment/giftDetail.png
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/215c6f5032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/2bfe252032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/0d79a52032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/1d5edbe032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/12ae801032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/17f8b7c032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/26e97c6032ee11e6a43b0f9416af7213.jpg
cid=5 : http://img5.bcyimg.com/coser/47840/post/17815/a3f10c30348f11e6b89e1967e57cfb1f.jpg
cid=5 : http://img5.bcyimg.com/coser/47840/post/17815/ac8f9f00348f11e6b89e1967e57cfb1f.jpg
cid=5 : http://img9.bcyimg.com/coser/47840/post/17815/7a7ae1a0348f11e6b89e1967e57cfb1f.jpg
cid=5 : http://img9.bcyimg.com/coser/47840/post/17815/852ee6f0348f11e6b89e1967e57cfb1f.jpg
cid=5 : http://img5.bcyimg.com/coser/47840/post/17815/8c4e2cc0348f11e6b89e1967e57cfb1f.jpg
cid=5 : http://img9.bcyimg.com/coser/47840/post/17815/939fa620348f11e6b89e1967e57cfb1f.jpg
cid=6 : http://img9.bcyimg.com/coser/47840/post/17815/98e4fbd0348f11e6b89e1967e57cfb1f.jpg
cid=7 : http://img5.bcyimg.com/coser/47840/post/17815/49563420349011e6aa3b1bfee02be998.jpg
cid=8 : http://img5.bcyimg.com/coser/47840/post/17813/1a60f67032f411e6941727176b80f3f3.jpg
cid=9 : http://img9.bcyimg.com/coser/47840/post/17813/e11d62e032f311e6941727176b80f3f3.jpg
cid=9 : http://img9.bcyimg.com/coser/47840/post/17813/246d4a1032f411e6941727176b80f3f3.jpg
cid=9 : http://img5.bcyimg.com/coser/47840/post/17813/ea0cfbe032f311e6941727176b80f3f3.jpg
cid=10 : http://img5.bcyimg.com/coser/47840/post/17818/c2cdcf1036e411e6b71051ef0ffbe5d8.jpg
cid=10 : http://img9.bcyimg.com/coser/47840/post/17818/d01aa7b036e411e6b71051ef0ffbe5d8.jpg
cid=10 : http://img9.bcyimg.com/coser/47840/post/17818/e0c2923036e411e6b71051ef0ffbe5d8.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/f269992032f211e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/f8b9a81032f211e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/fd80a60032f211e6941727176b80f3f3.jpg
cid=11 : http://img5.bcyimg.com/coser/47840/post/17813/0226876032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/0653412032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/0a88d48032f311e6941727176b80f3f3.jpg
cid=11 : http://img5.bcyimg.com/coser/47840/post/17813/0e07e92032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/1270e89032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/1660242032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/1a31295032f311e6941727176b80f3f3.jpg
cid=11 : http://img5.bcyimg.com/coser/47840/post/17813/1f32712032f311e6941727176b80f3f3.jpg
cid=11 : http://img5.bcyimg.com/coser/47840/post/17813/2337097032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/2801ff0032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/2cc9722032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/30fbd13032f311e6941727176b80f3f3.jpg
cid=11 : http://img5.bcyimg.com/coser/47840/post/17813/34446c8032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/390b918032f311e6941727176b80f3f3.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/17813/3dd4d96032f311e6941727176b80f3f3.jpg
cid=12 : http://img9.bcyimg.com/coser/47840/post/17818/f0241a0036e411e6b71051ef0ffbe5d8.jpg
cid=12 : http://img9.bcyimg.com/coser/47840/post/17818/f7b645e036e411e6b71051ef0ffbe5d8.jpg
cid=12 : http://img9.bcyimg.com/coser/47840/post/17818/0591c9f036e511e6b71051ef0ffbe5d8.jpg
cid=12 : http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg
cid=13 : /Public/Image/app.png
cid=14 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=15 : http://img9.bcyimg.com/coser/47840/post/4blp/8a4be4a027cc11e6883c1736fee6edeb.jpg
cid=16 : http://img5.bcyimg.com/coser/47840/post/17814/1687212033c411e6b1c1e563932fa8f9.jpg
cid=17 : http://img9.bcyimg.com/coser/47840/post/4bm2/22e159a041ec11e6a48aa1b424cea3b5.jpg
cid=18 : http://img9.bcyimg.com/coser/47840/post/17818/28a8bba036eb11e6b88433af95948698.jpg
cid=19 : http://img5.bcyimg.com/coser/47840/post/17814/1687212033c411e6b1c1e563932fa8f9.jpg
cid=20 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
cid=21 : http://img9.bcyimg.com/coser/47840/post/4blp/8a4be4a027cc11e6883c1736fee6edeb.jpg
cid=22 : http://img9.bcyimg.com/editor/flag/c04hy/0e964d00c80d11e690af7feef9a49a4a.jpg
cid=23 : http://img9.bcyimg.com/editor/flag/1789t/0f644a60862d11e6806365efe9e0d6f6.jpg
cid=24 : http://user.bcyimg.com/Public/Upload/avatar/1971889/8cb37ce0642c41609f0e98a3674aded8/middle.jpg
cid=25 : http://user.bcyimg.com/Public/Upload/avatar/1971495/eefef6831d414dd1aa02134788b01497/middle.jpg
cid=26 : /Public/Image/user_pic_middle.png
cid=27 : http://user.bcyimg.com/Public/Upload/avatar/829981/d345cba216ec43de96ec3a46967585b9/middle.jpg
cid=27 : http://user.bcyimg.com/Public/Upload/avatar/829981/d345cba216ec43de96ec3a46967585b9/middle.jpg
cid=28 : /Public/Image/user_pic_middle.png
cid=29 : http://user.bcyimg.com/Public/Upload/avatar/1722841/5771c24e17f4453392344f3b2d485273/middle.jpg
cid=30 : http://user.bcyimg.com/Public/Upload/avatar/1029245/42a6d27d3b0a4bb9aad95554d8bffabb/middle.jpg
cid=31 : /Public/Image/user_pic_middle.png
cid=32 : http://user.bcyimg.com/Public/Upload/avatar/1032988/f63d8ef8ac924dfd920a9c56f7c78563/middle.jpg
cid=33 : /Public/Image/user_pic_middle.png
cid=34 : http://user.bcyimg.com/Public/Upload/avatar/955015/29a05bd0218a4aa6a08a269c29a5daa7/middle.jpg
cid=35 : /Public/Image/user_pic_middle.png
cid=35 : /Public/Image/user_pic_middle.png
cid=36 : http://user.bcyimg.com/Public/Upload/avatar/1200649/feac93a618b84096a8d713ce65b86f78/middle.jpg
cid=37 : http://img9.bcyimg.com/article/post/c04hq/d6426890c1d511e6a6ecab0bcfeee30f.jpg
cid=38 : http://img9.bcyimg.com/article/post/c04hy/eabc6500c7fc11e6b37e0f01456a9728.png
cid=39 : http://img9.bcyimg.com/article/post/c04hw/85897cc0c68c11e6be5951aaa87d60c3.jpg
cid=39 : http://img9.bcyimg.com/article/post/c04hw/6abd22c0c68c11e6be5951aaa87d60c3.jpg
cid=40 : http://img9.bcyimg.com/article/post/c04hv/db2f9230c5b111e6b83479192b28650f.jpg
cid=41 : /Public/Image/app.png
cid=42 : /Public/Image/qr-cltAndr.gif
```

I want download those links.

* cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
* ...
* cid=12 : http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg

those links have different cluster id (cid).
use `-similarity=0.85` to make new cluster. (default similarity is 0.9)

```
./rise -cmd=show -uri=http://bcy.net/coser/detail/47840/602658 -filter=1 -similarity=0.85

cid=0 : /Public/Image/zhipin/tip-zhipin.png
cid=1 : /Public/Image/sub-nav/logo-cos-channel.png
cid=2 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=2 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=3 : /Public/Image/payment/giftDetail.png
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/215c6f5032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/2bfe252032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/0d79a52032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/1d5edbe032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/12ae801032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/17f8b7c032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/26e97c6032ee11e6a43b0f9416af7213.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17815/a3f10c30348f11e6b89e1967e57cfb1f.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17815/ac8f9f00348f11e6b89e1967e57cfb1f.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17815/7a7ae1a0348f11e6b89e1967e57cfb1f.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17815/852ee6f0348f11e6b89e1967e57cfb1f.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17815/8c4e2cc0348f11e6b89e1967e57cfb1f.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17815/939fa620348f11e6b89e1967e57cfb1f.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17815/98e4fbd0348f11e6b89e1967e57cfb1f.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17815/49563420349011e6aa3b1bfee02be998.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/1a60f67032f411e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/e11d62e032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/246d4a1032f411e6941727176b80f3f3.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/ea0cfbe032f311e6941727176b80f3f3.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17818/c2cdcf1036e411e6b71051ef0ffbe5d8.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17818/d01aa7b036e411e6b71051ef0ffbe5d8.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17818/e0c2923036e411e6b71051ef0ffbe5d8.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/f269992032f211e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/f8b9a81032f211e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/fd80a60032f211e6941727176b80f3f3.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/0226876032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/0653412032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/0a88d48032f311e6941727176b80f3f3.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/0e07e92032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/1270e89032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/1660242032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/1a31295032f311e6941727176b80f3f3.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/1f32712032f311e6941727176b80f3f3.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/2337097032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/2801ff0032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/2cc9722032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/30fbd13032f311e6941727176b80f3f3.jpg
cid=4 : http://img5.bcyimg.com/coser/47840/post/17813/34446c8032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/390b918032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/3dd4d96032f311e6941727176b80f3f3.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17818/f0241a0036e411e6b71051ef0ffbe5d8.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17818/f7b645e036e411e6b71051ef0ffbe5d8.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17818/0591c9f036e511e6b71051ef0ffbe5d8.jpg
cid=4 : http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg
cid=5 : /Public/Image/app.png
cid=6 : http://user.bcyimg.com/Public/Upload/avatar/471323/a37052e43a4c45c7a17c250ede8f8701/big.jpg
cid=7 : http://img9.bcyimg.com/coser/47840/post/4blp/8a4be4a027cc11e6883c1736fee6edeb.jpg
cid=8 : http://img5.bcyimg.com/coser/47840/post/17814/1687212033c411e6b1c1e563932fa8f9.jpg
cid=9 : http://img9.bcyimg.com/coser/47840/post/4bm2/22e159a041ec11e6a48aa1b424cea3b5.jpg
cid=10 : http://img9.bcyimg.com/coser/47840/post/17818/28a8bba036eb11e6b88433af95948698.jpg
cid=10 : http://img5.bcyimg.com/coser/47840/post/17814/1687212033c411e6b1c1e563932fa8f9.jpg
cid=10 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
cid=11 : http://img9.bcyimg.com/coser/47840/post/4blp/8a4be4a027cc11e6883c1736fee6edeb.jpg
cid=12 : http://img9.bcyimg.com/editor/flag/c04hy/0e964d00c80d11e690af7feef9a49a4a.jpg
cid=13 : http://img9.bcyimg.com/editor/flag/1789t/0f644a60862d11e6806365efe9e0d6f6.jpg
cid=14 : http://user.bcyimg.com/Public/Upload/avatar/1971889/8cb37ce0642c41609f0e98a3674aded8/middle.jpg
cid=14 : http://user.bcyimg.com/Public/Upload/avatar/1971495/eefef6831d414dd1aa02134788b01497/middle.jpg
cid=15 : /Public/Image/user_pic_middle.png
cid=16 : http://user.bcyimg.com/Public/Upload/avatar/829981/d345cba216ec43de96ec3a46967585b9/middle.jpg
cid=16 : http://user.bcyimg.com/Public/Upload/avatar/829981/d345cba216ec43de96ec3a46967585b9/middle.jpg
cid=17 : /Public/Image/user_pic_middle.png
cid=18 : http://user.bcyimg.com/Public/Upload/avatar/1722841/5771c24e17f4453392344f3b2d485273/middle.jpg
cid=19 : http://user.bcyimg.com/Public/Upload/avatar/1029245/42a6d27d3b0a4bb9aad95554d8bffabb/middle.jpg
cid=20 : /Public/Image/user_pic_middle.png
cid=21 : http://user.bcyimg.com/Public/Upload/avatar/1032988/f63d8ef8ac924dfd920a9c56f7c78563/middle.jpg
cid=22 : /Public/Image/user_pic_middle.png
cid=23 : http://user.bcyimg.com/Public/Upload/avatar/955015/29a05bd0218a4aa6a08a269c29a5daa7/middle.jpg
cid=24 : /Public/Image/user_pic_middle.png
cid=24 : /Public/Image/user_pic_middle.png
cid=25 : http://user.bcyimg.com/Public/Upload/avatar/1200649/feac93a618b84096a8d713ce65b86f78/middle.jpg
cid=26 : http://img9.bcyimg.com/article/post/c04hr/454f48e0c28d11e6ae4ce3c9637442d1.jpg
cid=26 : http://img9.bcyimg.com/article/post/c04hx/7da882a0c75211e69609e1e8c1b0018d.jpg
cid=26 : http://img9.bcyimg.com/article/post/c04hy/f33ce880c7fc11e6b37e0f01456a9728.jpg
cid=27 : http://img9.bcyimg.com/article/post/c04hx/87884800c75211e69609e1e8c1b0018d.jpg
cid=28 : http://img9.bcyimg.com/article/post/c04hw/7276acc0c68c11e6be5951aaa87d60c3.jpg
cid=29 : /Public/Image/app.png
cid=30 : /Public/Image/qr-cltAndr.gif
```

Links that I want to download have cluster id=4.

* cid=4 : http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
* ...
* cid=4 : http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg

use `-cid=4` to peek only cluster id=4.

```
./rise -cmd=show -uri=http://bcy.net/coser/detail/47840/602658 -filter=1 -similarity=0.85 -cid=4

(1/45) http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
(2/45) http://img9.bcyimg.com/coser/47840/post/17813/215c6f5032ee11e6a43b0f9416af7213.jpg
(3/45) http://img9.bcyimg.com/coser/47840/post/17813/2bfe252032ee11e6a43b0f9416af7213.jpg
(4/45) http://img5.bcyimg.com/coser/47840/post/17813/0d79a52032ee11e6a43b0f9416af7213.jpg
(5/45) http://img5.bcyimg.com/coser/47840/post/17813/1d5edbe032ee11e6a43b0f9416af7213.jpg
(6/45) http://img5.bcyimg.com/coser/47840/post/17813/12ae801032ee11e6a43b0f9416af7213.jpg
(7/45) http://img5.bcyimg.com/coser/47840/post/17813/17f8b7c032ee11e6a43b0f9416af7213.jpg
(8/45) http://img9.bcyimg.com/coser/47840/post/17813/26e97c6032ee11e6a43b0f9416af7213.jpg
(9/45) http://img5.bcyimg.com/coser/47840/post/17815/a3f10c30348f11e6b89e1967e57cfb1f.jpg
(10/45) http://img5.bcyimg.com/coser/47840/post/17815/ac8f9f00348f11e6b89e1967e57cfb1f.jpg
(11/45) http://img9.bcyimg.com/coser/47840/post/17815/7a7ae1a0348f11e6b89e1967e57cfb1f.jpg
(12/45) http://img9.bcyimg.com/coser/47840/post/17815/852ee6f0348f11e6b89e1967e57cfb1f.jpg
(13/45) http://img5.bcyimg.com/coser/47840/post/17815/8c4e2cc0348f11e6b89e1967e57cfb1f.jpg
(14/45) http://img9.bcyimg.com/coser/47840/post/17815/939fa620348f11e6b89e1967e57cfb1f.jpg
(15/45) http://img9.bcyimg.com/coser/47840/post/17815/98e4fbd0348f11e6b89e1967e57cfb1f.jpg
(16/45) http://img5.bcyimg.com/coser/47840/post/17815/49563420349011e6aa3b1bfee02be998.jpg
(17/45) http://img5.bcyimg.com/coser/47840/post/17813/1a60f67032f411e6941727176b80f3f3.jpg
(18/45) http://img9.bcyimg.com/coser/47840/post/17813/e11d62e032f311e6941727176b80f3f3.jpg
(19/45) http://img9.bcyimg.com/coser/47840/post/17813/246d4a1032f411e6941727176b80f3f3.jpg
(20/45) http://img5.bcyimg.com/coser/47840/post/17813/ea0cfbe032f311e6941727176b80f3f3.jpg
(21/45) http://img5.bcyimg.com/coser/47840/post/17818/c2cdcf1036e411e6b71051ef0ffbe5d8.jpg
(22/45) http://img9.bcyimg.com/coser/47840/post/17818/d01aa7b036e411e6b71051ef0ffbe5d8.jpg
(23/45) http://img9.bcyimg.com/coser/47840/post/17818/e0c2923036e411e6b71051ef0ffbe5d8.jpg
(24/45) http://img9.bcyimg.com/coser/47840/post/17813/f269992032f211e6941727176b80f3f3.jpg
(25/45) http://img9.bcyimg.com/coser/47840/post/17813/f8b9a81032f211e6941727176b80f3f3.jpg
(26/45) http://img9.bcyimg.com/coser/47840/post/17813/fd80a60032f211e6941727176b80f3f3.jpg
(27/45) http://img5.bcyimg.com/coser/47840/post/17813/0226876032f311e6941727176b80f3f3.jpg
(28/45) http://img9.bcyimg.com/coser/47840/post/17813/0653412032f311e6941727176b80f3f3.jpg
(29/45) http://img9.bcyimg.com/coser/47840/post/17813/0a88d48032f311e6941727176b80f3f3.jpg
(30/45) http://img5.bcyimg.com/coser/47840/post/17813/0e07e92032f311e6941727176b80f3f3.jpg
(31/45) http://img9.bcyimg.com/coser/47840/post/17813/1270e89032f311e6941727176b80f3f3.jpg
(32/45) http://img9.bcyimg.com/coser/47840/post/17813/1660242032f311e6941727176b80f3f3.jpg
(33/45) http://img9.bcyimg.com/coser/47840/post/17813/1a31295032f311e6941727176b80f3f3.jpg
(34/45) http://img5.bcyimg.com/coser/47840/post/17813/1f32712032f311e6941727176b80f3f3.jpg
(35/45) http://img5.bcyimg.com/coser/47840/post/17813/2337097032f311e6941727176b80f3f3.jpg
(36/45) http://img9.bcyimg.com/coser/47840/post/17813/2801ff0032f311e6941727176b80f3f3.jpg
(37/45) http://img9.bcyimg.com/coser/47840/post/17813/2cc9722032f311e6941727176b80f3f3.jpg
(38/45) http://img9.bcyimg.com/coser/47840/post/17813/30fbd13032f311e6941727176b80f3f3.jpg
(39/45) http://img5.bcyimg.com/coser/47840/post/17813/34446c8032f311e6941727176b80f3f3.jpg
(40/45) http://img9.bcyimg.com/coser/47840/post/17813/390b918032f311e6941727176b80f3f3.jpg
(41/45) http://img9.bcyimg.com/coser/47840/post/17813/3dd4d96032f311e6941727176b80f3f3.jpg
(42/45) http://img9.bcyimg.com/coser/47840/post/17818/f0241a0036e411e6b71051ef0ffbe5d8.jpg
(43/45) http://img9.bcyimg.com/coser/47840/post/17818/f7b645e036e411e6b71051ef0ffbe5d8.jpg
(44/45) http://img9.bcyimg.com/coser/47840/post/17818/0591c9f036e511e6b71051ef0ffbe5d8.jpg
(45/45) http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg
```

I will download 45 links. use `-cmd=fetch` and `-output=kasugano-sora.zip` to save images.
If you `-worker=16`, you can use multiple workers.

```
./rise -cmd=fetch -uri=http://bcy.net/coser/detail/47840/602658 -filter=1 -similarity=0.85 -cid=4 -output=kasugano-sora.zip -worker=16

2016/12/24 17:49:34 [worker 14] download http://img5.bcyimg.com/coser/47840/post/17815/8c4e2cc0348f11e6b89e1967e57cfb1f.jpg
2016/12/24 17:49:34 [worker 4] download http://img9.bcyimg.com/coser/47840/post/17815/939fa620348f11e6b89e1967e57cfb1f.jpg
2016/12/24 17:49:34 [worker 7] download http://img5.bcyimg.com/coser/47840/post/17813/1d5edbe032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 15] download http://img9.bcyimg.com/coser/47840/post/17815/98e4fbd0348f11e6b89e1967e57cfb1f.jpg
2016/12/24 17:49:34 [worker 16] download http://img5.bcyimg.com/coser/47840/post/17815/49563420349011e6aa3b1bfee02be998.jpg
2016/12/24 17:49:34 [worker 2] download http://img9.bcyimg.com/coser/47840/post/17813/083f4bf032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 1] download http://img9.bcyimg.com/coser/47840/post/17813/215c6f5032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 6] download http://img9.bcyimg.com/coser/47840/post/17813/2bfe252032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 5] download http://img5.bcyimg.com/coser/47840/post/17813/0d79a52032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 10] download http://img5.bcyimg.com/coser/47840/post/17815/a3f10c30348f11e6b89e1967e57cfb1f.jpg
2016/12/24 17:49:34 [worker 3] download http://img5.bcyimg.com/coser/47840/post/17813/12ae801032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 8] download http://img5.bcyimg.com/coser/47840/post/17813/17f8b7c032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 9] download http://img9.bcyimg.com/coser/47840/post/17813/26e97c6032ee11e6a43b0f9416af7213.jpg
2016/12/24 17:49:34 [worker 12] download http://img9.bcyimg.com/coser/47840/post/17815/7a7ae1a0348f11e6b89e1967e57cfb1f.jpg
2016/12/24 17:49:34 [worker 11] download http://img5.bcyimg.com/coser/47840/post/17815/ac8f9f00348f11e6b89e1967e57cfb1f.jpg
2016/12/24 17:49:34 [worker 16] download http://img5.bcyimg.com/coser/47840/post/17818/c2cdcf1036e411e6b71051ef0ffbe5d8.jpg
2016/12/24 17:49:34 [worker 10] download http://img9.bcyimg.com/coser/47840/post/17818/e0c2923036e411e6b71051ef0ffbe5d8.jpg
2016/12/24 17:49:34 [worker 15] download http://img9.bcyimg.com/coser/47840/post/17813/f8b9a81032f211e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 4] download http://img5.bcyimg.com/coser/47840/post/17813/1a60f67032f411e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 12] download http://img9.bcyimg.com/coser/47840/post/17813/0a88d48032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 7] download http://img9.bcyimg.com/coser/47840/post/17813/246d4a1032f411e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 14] download http://img5.bcyimg.com/coser/47840/post/17813/ea0cfbe032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 complete (1/45)
2016/12/24 17:49:34 [worker 16] download http://img9.bcyimg.com/coser/47840/post/17813/1660242032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 4] download http://img9.bcyimg.com/coser/47840/post/17813/1a31295032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 13] download http://img9.bcyimg.com/coser/47840/post/17815/852ee6f0348f11e6b89e1967e57cfb1f.jpg
2016/12/24 17:49:34 [worker 11] download http://img5.bcyimg.com/coser/47840/post/17813/1f32712032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 12] download http://img9.bcyimg.com/coser/47840/post/17813/2cc9722032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 8] download http://img9.bcyimg.com/coser/47840/post/17813/f269992032f211e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 6] download http://img9.bcyimg.com/coser/47840/post/17813/e11d62e032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 3] download http://img9.bcyimg.com/coser/47840/post/17813/fd80a60032f211e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 1] download http://img5.bcyimg.com/coser/47840/post/17813/0226876032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 9] download http://img9.bcyimg.com/coser/47840/post/17813/0653412032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 15] download http://img5.bcyimg.com/coser/47840/post/17813/0e07e92032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 5] download http://img9.bcyimg.com/coser/47840/post/17813/1270e89032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 4] download http://img5.bcyimg.com/coser/47840/post/17813/2337097032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 2] download http://img9.bcyimg.com/coser/47840/post/17818/d01aa7b036e411e6b71051ef0ffbe5d8.jpg
2016/12/24 17:49:34 [worker 13] download http://img9.bcyimg.com/coser/47840/post/17813/2801ff0032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 16] download http://img9.bcyimg.com/coser/47840/post/17813/30fbd13032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 8] download http://img5.bcyimg.com/coser/47840/post/17813/34446c8032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 12] download http://img9.bcyimg.com/coser/47840/post/17813/390b918032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 14] download http://img9.bcyimg.com/coser/47840/post/17813/3dd4d96032f311e6941727176b80f3f3.jpg
2016/12/24 17:49:34 [worker 10] download http://img9.bcyimg.com/coser/47840/post/17818/f0241a0036e411e6b71051ef0ffbe5d8.jpg
2016/12/24 17:49:34 [worker 1] download http://img9.bcyimg.com/coser/47840/post/17818/f7b645e036e411e6b71051ef0ffbe5d8.jpg
2016/12/24 17:49:34 [worker 6] download http://img9.bcyimg.com/coser/47840/post/17818/0591c9f036e511e6b71051ef0ffbe5d8.jpg
2016/12/24 17:49:34 [worker 11] download http://img9.bcyimg.com/coser/47840/post/17818/0d23cec036e511e6b71051ef0ffbe5d8.jpg
2016/12/24 17:49:34 complete (2/45)
2016/12/24 17:49:34 complete (3/45)
2016/12/24 17:49:34 complete (4/45)
2016/12/24 17:49:34 complete (5/45)
2016/12/24 17:49:34 complete (6/45)
2016/12/24 17:49:34 complete (7/45)
2016/12/24 17:49:35 complete (8/45)
2016/12/24 17:49:35 complete (9/45)
2016/12/24 17:49:35 complete (10/45)
2016/12/24 17:49:35 complete (11/45)
2016/12/24 17:49:35 complete (12/45)
2016/12/24 17:49:35 complete (13/45)
2016/12/24 17:49:35 complete (14/45)
2016/12/24 17:49:35 complete (15/45)
2016/12/24 17:49:35 complete (16/45)
2016/12/24 17:49:36 complete (17/45)
2016/12/24 17:49:36 complete (18/45)
2016/12/24 17:49:36 complete (19/45)
2016/12/24 17:49:36 complete (20/45)
2016/12/24 17:49:36 complete (21/45)
2016/12/24 17:49:36 complete (22/45)
2016/12/24 17:49:36 complete (23/45)
2016/12/24 17:49:36 complete (24/45)
2016/12/24 17:49:36 complete (25/45)
2016/12/24 17:49:36 complete (26/45)
2016/12/24 17:49:37 complete (27/45)
2016/12/24 17:49:37 complete (28/45)
2016/12/24 17:49:37 complete (29/45)
2016/12/24 17:49:37 complete (30/45)
2016/12/24 17:49:37 complete (31/45)
2016/12/24 17:49:37 complete (32/45)
2016/12/24 17:49:37 complete (33/45)
2016/12/24 17:49:37 complete (34/45)
2016/12/24 17:49:38 complete (35/45)
2016/12/24 17:49:38 complete (36/45)
2016/12/24 17:49:38 complete (37/45)
2016/12/24 17:49:38 complete (38/45)
2016/12/24 17:49:38 complete (39/45)
2016/12/24 17:49:38 complete (40/45)
2016/12/24 17:49:38 complete (41/45)
2016/12/24 17:49:38 complete (42/45)
2016/12/24 17:49:38 complete (43/45)
2016/12/24 17:49:39 complete (44/45)
2016/12/24 17:49:39 complete (45/45)
```

done!


## Options

* `-cmd=show` : show links
* `-cmd=fetch` : fetch images

* `-cid=1` : select cluster id

* `-simility=0.8` : similarity is used to create cluster. default similarity is 0.9.

* `-filter=0` : no filter. do not modify link. default value
* `-filter=1` : slice extra strings after extension
    * example: "http://google.com/helloworld.jpg/w650" -> "http://google.com/helloworld.jpg" (/w650 is sliced)

* `-worker=8` : download images with multiple workers.