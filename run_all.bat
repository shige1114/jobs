@echo off
setlocal enabledelayedexpansion

REM 00_dockerディレクトリ内のサブディレクトリをループ
for /d %%d in (00_docker\*) do (
    REM サブディレクトリにdocker-compose.ymlが存在するか確認
    if exist "%%d\docker-compose.yml" (
        echo Running docker-compose in %%d
        REM サブディレクトリに移動してdocker-compose up -dを実行
        pushd %%d
        docker-compose up -d
        popd
    ) else (
        echo No docker-compose.yml found in %%d
    )
)

endlocal
pause

