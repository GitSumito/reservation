
package config

import (
    "github.com/sirupsen/logrus"
)

const (
    // MySQLHost ...
    MySQLHost = "reservation_mysql_1"
    // MySQLUser ...
    MySQLUser = "hayashi"
    // MySQLPassword ...
    MySQLPassword = "1234"
    // MySQLPort ...
    MySQLPort = "3306"
    // MySQLTimeZone ...
    MySQLTimeZone = "Asia%2FTokyo"

    //GinMode debug/release/test

    // LogLevel ...
    LogLevel = logrus.DebugLevel
    // LogFile ...
    LogFile = "/go/src/api-sample/logs/application.log"
)