#!/usr/bin/python

import pycurl
import json

c = pycurl.Curl()

data = {"method": "TRHService.TRH", "params": [{"Where:Jakarta", "Time":"2","Temp":"3","RH":"4"}], "id":"1"}

c.setopt(c.URL, 'http://10.10.5.42:10000/rpc')
c.setopt(c.HTTPHEADER, ["Content-Type: application/json"])
post = json.dumps(data)
c.setopt(c.POST, 1)
c.setopt(c.POSTFIELDS, post)

c.perform()