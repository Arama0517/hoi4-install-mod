Invoke-WebRequest -Uri https://partner.steamgames.com/downloads/steamworks_sdk.zip -OutFile .\sdk.zip
Expand-Archive -Path 'sdk.zip' -DestinationPath '.' -Force
Remove-Item -Path sdk.zip
Move-Item -Path sdk\redistributable_bin\steam_api.dll -Destination .
Move-Item -Path sdk\redistributable_bin\win64\steam_api64.dll -Destination .
Remove-Item -Path sdk -Recurse