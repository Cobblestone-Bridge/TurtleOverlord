local trId = tostring(os.getComputerID())
local shttp = http.get("http://localhost:8000/turtle/login/?id=" .. trId)
write(shttp.readAll())
shttp.close()
