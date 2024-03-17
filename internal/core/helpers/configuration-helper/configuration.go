package helpers

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	ServiceName               string `mapstructure:"Service_Name"`
	ServiceHost               string `mapstructure:"Service_Host"`
	ServicePort               string `mapstructure:"Service_Port"`
	MongoDBConnString         string `mapstructure:"MongoDB_Connection_String"`
	MongoDbDatabaseName       string `mapstructure:"MongoDB_Database_Name"`
	MongoDBUserCollectionName string `mapstructure:"MongoDB_Notification_Collection_Name"`
	RedisConnString           string `mapstructure:"Redis_Connection_String"`
	RedisConnPassword         string `mapstructure:"Redis_Connection_Password"`
	LogDir                    string `mapstructure:"Log_Dir"`
	LogFile                   string `mapstructure:"Log_File"`
	GoogleSmtpHost            string `mapstructure:"Google_Smtp_Host"`
	GoogleSmtpPort            string `mapstructure:"Google_Smtp_Port"`
	GoogleAppPassword         string `mapstructure:"Google_App_Password"`
	GoogleAuthUser            string `mapstructure:"Google_Auth_User"`
	TwilioAccountSID          string `mapstructure:"Twilio_Account_SID"`
	TwilioAuthToken           string `mapstructure:"Twilio_Auth_Token"`
	TwilioAuthPhoneNumber     string `mapstructure:"Twilio_Auth_Phone_Number"`
	MailgunDomain             string `mapstructure:"Mailgun_Domain"`
	MailgunPrivateKey         string `mapstructure:"Mailgun_Private_Key"`
	RedisRevokedTokensKey     string `mapstructure:"Redis_Revoked_Tokens_Key"`
	JWTTokenKey               string `mapstructure:"JWT_Token_Key"`
}

var ServiceConfiguration = loadConfig(".")

func loadConfig(path string) Configuration {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read in config:", err)
	}

	var config Configuration
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("unmarsal in config:", err)
	}

	return config
}
