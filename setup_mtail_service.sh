#! /bin/bash
sudo ln -s $(pwd)/mtail/mtail.service /etc/systemd/system/mtail.service
sudo systemctl daemon-reload
sudo systemctl restart mtail
