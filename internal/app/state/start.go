package state

import (
	"car_bot/internal/app/ds"
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type ChatContext struct {
	User *ds.User
	Vk   *api.VK
	Ctx  *context.Context
	Db   *sqlx.DB
}

type State interface {
	Name() string                                      //получаем название состояния в виде строки, чтобы в дальнейшем куда-то записать(БД)
	Process(ChatContext, object.MessagesMessage) State //нужно взять контекст, посмотреть на каком состоянии сейчас пользователь, метод должен вернуть состояние
	PreviewProcess(ctc ChatContext)
}

// ////////////////////////////////////////////////////////
type StartState struct {
}

func (state StartState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Начать тест" {
		_, err := ctc.Db.Exec("UPDATE users SET city_driving = $1, highway_driving = $2, fuel_efficient = $3, spacious = $4, powerful = $5, easy_to_drive = $6, comfortable = $7, sporty = $8, modern_design = $9, high_end_luxury = $10, environmentally_friendly = $11, fwd = $12, rwd = $13, awd = $14, rear_view_camera = $15, tire_pressure_monitoring = $16, automatic_parking = $17, navigation_system = $18, heated_seats = $19, electric_motor = $20  WHERE vk_id =$21", false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to reset user fields")
		}
		Question1{}.PreviewProcess(ctc)
		return &Question1{}
	} else {
		state.PreviewProcess(ctc)
		return &StartState{}
	}
}

func (state StartState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Тест: Какой автомобиль вам больше подходит? \nДля начала тестирования нажмите кнопку \"Начать тест\"")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Начать тест", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state StartState) Name() string {
	return "StartState"
}
