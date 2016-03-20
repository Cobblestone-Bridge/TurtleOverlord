-- The basic bootstrap.
-- Logs in and sends the computer ID and CraftOS version.
print( "[INFO] Starting boostrap program!" )

if not http then
  print( "[ERROR] The http library is not enabled!" )
  os.exit()
end

-- Information about computer.
local computerId = tostring( os.getComputerID() )
local computerVersion = tostring( os.version() )

-- Other variables.
local serverAddress = "localhost:8000"

-- Check if the computer is a turtle.
if turtle then
  print( "[INFO] This computer is a turtle!" )
end

 -- Connect function
local function connect( handle )
  write( "[INFO] Connecting to " .. serverAddress .. handle .. "... " )

  local ok, err = http.checkURL( serverAddress )
  if not ok then
    print( "[ERROR] Failed." )
    if err then
      printError( err )
    end
    return nil
  end

  local response = http.get( serverAddress .. handle )
  if not response then
    print( "[ERROR] Failed." )
    return nil
  end

  print( "[INFO] Success." )

  local sResponse = response.readAll()
  response.close()
  return sResponse
end

-- Check if md5 library is installed.
local sPath = shell.resolve( "TurtleOverlord/md5.lua" )
if not fs.exists( sPath ) then
  print ("[INFO] Downloading md5 library...")
  local response = connect( "/static/md5.lua" )
  if response then
    local file = fs.open( sPath, "w" )
    file.write( res )
    file.close()

    -- Load md5 library
    require("TurtleOverlord/md5.lua")
  end
end

-- Login to the server.
print("[INFO] Logging in...")
local response = connect( "/bootstrap/?id=" .. computerId .. "&version=" .. computerVersion )
if response then
  
end

-- Open the driver file.
local sPath = shell.resolve( sFile )
if fs.exists( sPath ) then
    print( "File already exists" )
    return
end
