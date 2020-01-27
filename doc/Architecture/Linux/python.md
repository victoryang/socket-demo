# Python

[pipreqs](https://blog.csdn.net/weixin_41173374/article/details/100623179)

[virtualenv](https://virtualenv.pypa.io/en/latest/userguide/)

[python3 env](https://blog.csdn.net/daily_to_full/article/details/89042278)


## Graphite

[installation](https://blog.csdn.net/weixin_39922154/article/details/79132147)

### Installation

[dependency](https://graphite.readthedocs.io/en/latest/install.html#dependencies)

virtualenv方式安装
```
// system-wide dependency

// virtualenv
apt install -y virtualenv
virtualenv /opt/graphite/
source /opt/graphite/bin/activate

// installed from source
// Does not recommand 'pip install -r reqirements'

pip install gunicorn
pip install django
pip install django-tagging
pip install scandir
pip install pyparsing
pip install cairocffi
pip install whisper

// for latest graphite-web
pip install six
pip install requests

mkdir -p storage/{ceres,whisper,log/webapp}

// 静态目录
PYTHONPATH=/opt/graphite/webapp django-admin.py collectstatic --noinput --settings=graphite.settings

// 数据库
PYTHONPATH=/opt/graphite/webapp django-admin.py migrate --settings=graphite.settings

chown nobody:nogroup /opt/graphite/storage/graphite.db

// edit local_settings.py
// TZ
TIME_ZONE = 'Asia/Shanghai'
// carbon
STANDARD_DIRS = ["/bigd/graphite/whisper/3"]

// run
export PYTHONPATH=/opt/graphite/webapp
/opt/graphite/bin/gunicorn -b 0.0.0.0:8090 -w 16 graphite.wsgi:application
```