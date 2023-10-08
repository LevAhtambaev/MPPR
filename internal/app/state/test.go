package state

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type Question1 struct {
}

func (state Question1) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET city_driving = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question2{}.PreviewProcess(ctc)
		return &Question2{}
	} else if messageText == "Нет" {
		Question2{}.PreviewProcess(ctc)
		return &Question2{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question1{}
	}
}

func (state Question1) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("1. Вам нужен автомобиль для городской езды?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question1) Name() string {
	return "Question1"
}

// ////////////////////////////////////////////////////////
type Question2 struct {
}

func (state Question2) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET highway_driving = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question3{}.PreviewProcess(ctc)
		return &Question3{}
	} else if messageText == "Нет" {
		Question3{}.PreviewProcess(ctc)
		return &Question3{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question2{}
	}
}

func (state Question2) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("2. Вам нужен автомобиль для дальних поездок?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question2) Name() string {
	return "Question2"
}

// ////////////////////////////////////////////////////////
type Question3 struct {
}

func (state Question3) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET spacious = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question4{}.PreviewProcess(ctc)
		return &Question4{}
	} else if messageText == "Нет" {
		Question4{}.PreviewProcess(ctc)
		return &Question4{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question3{}
	}
}

func (state Question3) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("3. Вам нужен автомобиль с большим багажником?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question3) Name() string {
	return "Question3"
}

// ////////////////////////////////////////////////////////
type Question4 struct {
}

func (state Question4) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET fuel_efficient = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question5{}.PreviewProcess(ctc)
		return &Question5{}
	} else if messageText == "Нет" {
		Question5{}.PreviewProcess(ctc)
		return &Question5{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question4{}
	}
}

func (state Question4) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("4. Вам нужен автомобиль с низким расходом топлива?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question4) Name() string {
	return "Question4"
}

// ////////////////////////////////////////////////////////
type Question5 struct {
}

func (state Question5) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET powerful = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question6{}.PreviewProcess(ctc)
		return &Question6{}
	} else if messageText == "Нет" {
		Question6{}.PreviewProcess(ctc)
		return &Question6{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question5{}
	}
}

func (state Question5) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("5. Вам нужен автомобиль с большой мощностью двигателя?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question5) Name() string {
	return "Question5"
}

// ////////////////////////////////////////////////////////
type Question6 struct {
}

func (state Question6) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET environmentally_friendly = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question7{}.PreviewProcess(ctc)
		return &Question7{}
	} else if messageText == "Нет" {
		Question7{}.PreviewProcess(ctc)
		return &Question7{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question6{}
	}
}

func (state Question6) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("6. Вам нужен автомобиль с низким уровнем выбросов?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question6) Name() string {
	return "Question6"
}

// ////////////////////////////////////////////////////////
type Question7 struct {
}

func (state Question7) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET easy_to_drive = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question8{}.PreviewProcess(ctc)
		return &Question8{}
	} else if messageText == "Нет" {
		Question8{}.PreviewProcess(ctc)
		return &Question8{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question7{}
	}
}

func (state Question7) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("7. Вам нужен легкий в управлении автомобиль?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question7) Name() string {
	return "Question7"
}

// ////////////////////////////////////////////////////////
type Question8 struct {
}

func (state Question8) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET comfortable = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question9{}.PreviewProcess(ctc)
		return &Question9{}
	} else if messageText == "Нет" {
		Question9{}.PreviewProcess(ctc)
		return &Question9{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question8{}
	}
}

func (state Question8) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("8. Вам нужен автомобиль с комфортным салоном?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question8) Name() string {
	return "Question8"
}

// ////////////////////////////////////////////////////////
type Question9 struct {
}

func (state Question9) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET sporty = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question10{}.PreviewProcess(ctc)
		return &Question10{}
	} else if messageText == "Нет" {
		Question10{}.PreviewProcess(ctc)
		return &Question10{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question9{}
	}
}

func (state Question9) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("9. Вам нужен спортивный автомобиль?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question9) Name() string {
	return "Question9"
}

// ////////////////////////////////////////////////////////
type Question10 struct {
}

func (state Question10) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET modern_design = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question11{}.PreviewProcess(ctc)
		return &Question11{}
	} else if messageText == "Нет" {
		Question11{}.PreviewProcess(ctc)
		return &Question11{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question10{}
	}
}

func (state Question10) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("10. Вам нужен с современным дизайном?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question10) Name() string {
	return "Question10"
}

// ////////////////////////////////////////////////////////
type Question11 struct {
}

func (state Question11) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET high_end_luxury = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question12{}.PreviewProcess(ctc)
		return &Question12{}
	} else if messageText == "Нет" {
		Question12{}.PreviewProcess(ctc)
		return &Question12{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question11{}
	}
}

func (state Question11) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("11. Вам нужен роскошный автомобиль?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question11) Name() string {
	return "Question11"
}

// ////////////////////////////////////////////////////////
type Question12 struct {
}

func (state Question12) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET fwd = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question13{}.PreviewProcess(ctc)
		return &Question13{}
	} else if messageText == "Нет" {
		Question13{}.PreviewProcess(ctc)
		return &Question13{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question12{}
	}
}

func (state Question12) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("12. Вам нужен автомобиль с передним приводом?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question12) Name() string {
	return "Question12"
}

// ////////////////////////////////////////////////////////
type Question13 struct {
}

func (state Question13) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET rwd = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question14{}.PreviewProcess(ctc)
		return &Question14{}
	} else if messageText == "Нет" {
		Question14{}.PreviewProcess(ctc)
		return &Question14{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question13{}
	}
}

func (state Question13) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("13. Вам нужен автомобиль с задним приводом?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question13) Name() string {
	return "Question13"
}

// ////////////////////////////////////////////////////////
type Question14 struct {
}

func (state Question14) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET awd = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question15{}.PreviewProcess(ctc)
		return &Question15{}
	} else if messageText == "Нет" {
		Question15{}.PreviewProcess(ctc)
		return &Question15{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question14{}
	}
}

func (state Question14) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("14. Вам нужен автомобиль с полным приводом?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question14) Name() string {
	return "Question14"
}

// ////////////////////////////////////////////////////////
type Question15 struct {
}

func (state Question15) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET rear_view_camera = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question16{}.PreviewProcess(ctc)
		return &Question16{}
	} else if messageText == "Нет" {
		Question16{}.PreviewProcess(ctc)
		return &Question16{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question15{}
	}
}

func (state Question15) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("15. Вам нужен автомобиль с камерой заднего вида?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question15) Name() string {
	return "Question15"
}

// ////////////////////////////////////////////////////////
type Question16 struct {
}

func (state Question16) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET tire_pressure_monitoring = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question17{}.PreviewProcess(ctc)
		return &Question17{}
	} else if messageText == "Нет" {
		Question17{}.PreviewProcess(ctc)
		return &Question17{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question16{}
	}
}

func (state Question16) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("16. Вам нужен автомобиль c контролем давления в шинах?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question16) Name() string {
	return "Question16"
}

// ////////////////////////////////////////////////////////
type Question17 struct {
}

func (state Question17) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET automatic_parking = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question18{}.PreviewProcess(ctc)
		return &Question18{}
	} else if messageText == "Нет" {
		Question18{}.PreviewProcess(ctc)
		return &Question18{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question17{}
	}
}

func (state Question17) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("17. Вам нужен автомобиль c автоматической парковкой?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question17) Name() string {
	return "Question17"
}

// ////////////////////////////////////////////////////////
type Question18 struct {
}

func (state Question18) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET navigation_system = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question19{}.PreviewProcess(ctc)
		return &Question19{}
	} else if messageText == "Нет" {
		Question19{}.PreviewProcess(ctc)
		return &Question19{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question18{}
	}
}

func (state Question18) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("18. Вам нужен автомобиль с навигационной системой?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question18) Name() string {
	return "Question18"
}

// ////////////////////////////////////////////////////////
type Question19 struct {
}

func (state Question19) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET heated_seats = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question20{}.PreviewProcess(ctc)
		return &Question20{}
	} else if messageText == "Нет" {
		Question20{}.PreviewProcess(ctc)
		return &Question20{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question19{}
	}
}

func (state Question19) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("19. Вам нужен автомобиль с подогревом сидений?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question19) Name() string {
	return "Question19"
}

// ////////////////////////////////////////////////////////
type Question20 struct {
}

func (state Question20) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET electric_motor = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		ResultsState{}.PreviewProcess(ctc)
		return &ResultsState{}
	} else if messageText == "Нет" {
		ResultsState{}.PreviewProcess(ctc)
		return &ResultsState{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question20{}
	}
}

func (state Question20) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("20. Вам нужен электромобиль?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question20) Name() string {
	return "Question20"
}
