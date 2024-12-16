package main

import (
	"log"
	"os"
	"telegram_bot/api_service"
	"telegram_bot/data_representation"
	"telegram_bot/stages"
	"telegram_bot/user_info"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"errors"
)

func main() {
	telegramApiKey := os.Getenv("TELEGRAM_API_KEY")
	bot, err := tgbotapi.NewBotAPI(telegramApiKey)
	if err != nil {
		log.Panic(err)
		return
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	userInfo := user_info.NewUserInfo()

	apiService, errService := api_service.NewService()
	if errService != nil {
		log.Panic(errService)
		return
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добрый день")

		currUser := userInfo.GetUserData(update.Message.From.ID)

		if update.Message.IsCommand() {
			if update.Message.Command() == "start" {
				initialStage := stages.NewInitialStage()
				userInfo.SetUserStage(update.Message.From.ID, initialStage)
				initialStage.ConfigureMessage(&msg)
			}
		} else {
			if userInfo.UserStageExists(update.Message.From.ID) {
				userStage := userInfo.GetUserStage(update.Message.From.ID)

				if _, resultInitialStage := userStage.(*stages.InitialStage); resultInitialStage {
					switch update.Message.Text {
					case "Войти":
						loginStage := stages.NewLoginStage()
						userInfo.SetUserStage(update.Message.From.ID, loginStage)
						loginStage.ConfigureMessage(&msg)
					case "Зарегистрироваться":
						registerStage := stages.NewRegisterStage()
						userInfo.SetUserStage(update.Message.From.ID, registerStage)
						registerStage.ConfigureMessage(&msg)
					default:
						msg.Text = "Не знаю такой команды"
					}
				}

				if registerStage, resultRegisterStage := userStage.(*stages.RegisterStage); resultRegisterStage {
					inputError := registerStage.ComputeInput(userInfo, update.Message.From.ID, update.Message.Text)
					if inputError != nil {
						newInitialStage := stages.NewInitialStage()
						userInfo.SetUserStage(update.Message.From.ID, newInitialStage)
						newInitialStage.ConfigureMessage(&msg)
						msg.Text = "Ошибка при вводе данных. " + msg.Text
					} else {
						registerStage.EndStage()
						if registerStage.StagesFinished() {
							newInitialStage := stages.NewInitialStage()
							userInfo.SetUserStage(update.Message.From.ID, newInitialStage)
							newInitialStage.ConfigureMessage(&msg)
							registerResult, errRegister := registerStage.Finish(currUser, apiService)
							if errRegister != nil {
								msg.Text = "Ошибка регистрации. " + msg.Text
							} else {
								if registerResult {
									msg.Text = "Регистрация успешна. " + msg.Text
								} else {
									msg.Text = "Регистрация не удалась. " + msg.Text
								}
							}
						} else {
							registerStage.ConfigureMessage(&msg)
						}
					}
				}

				if loginStage, resultLoginStage := userStage.(*stages.LoginStage); resultLoginStage {
					inputError := loginStage.ComputeInput(userInfo, update.Message.From.ID, update.Message.Text)
					if inputError != nil {
						newInitialStage := stages.NewInitialStage()
						userInfo.SetUserStage(update.Message.From.ID, newInitialStage)
						newInitialStage.ConfigureMessage(&msg)
						msg.Text = "Ошибка при вводе данных. " + msg.Text
					} else {
						loginStage.EndStage()
						if loginStage.StagesFinished() {
							loginResult, errLogin := loginStage.Finish(currUser, apiService)

							if loginResult {
								newMainStage := stages.NewMainStage()
								userInfo.SetUserStage(update.Message.From.ID, newMainStage)
								newMainStage.ConfigureMessage(&msg)
								msg.Text = "Вход успешен. " + msg.Text
							} else {
								newInitialStage := stages.NewInitialStage()
								userInfo.SetUserStage(update.Message.From.ID, newInitialStage)
								newInitialStage.ConfigureMessage(&msg)
								if errLogin != nil {
									msg.Text = "Ошибка при входе. " + msg.Text
								} else {
									msg.Text = "Вход не удался. " + msg.Text
								}
							}
						} else {
							loginStage.ConfigureMessage(&msg)
						}
					}
				}

				if _, resultMainStage := userStage.(*stages.MainStage); resultMainStage {
					if currUser.Role == nil {
						log.Fatal(errors.New("role is nil"))
					}
					switch update.Message.Text {
					case "Отели":
						hotelMenuStage := stages.NewHotelsMenuStage(*currUser.Role)
						userInfo.SetUserStage(update.Message.From.ID, hotelMenuStage)
						hotelMenuStage.ConfigureMessage(&msg)
					case "Бронирования":
						bookingMenuStage := stages.NewBookingsMenuStage(*currUser.Role)
						userInfo.SetUserStage(update.Message.From.ID, bookingMenuStage)
						bookingMenuStage.ConfigureMessage(&msg)
					case "Выйти":
						errLogout := apiService.Logout(currUser)
						if errLogout != nil {
							msg.Text = "Ошибка при выходе"
						} else {
							newInitialStage := stages.NewInitialStage()
							userInfo.SetUserStage(update.Message.From.ID, newInitialStage)
							newInitialStage.ConfigureMessage(&msg)
							msg.Text = "Вы успешно вышли. " + msg.Text
						}
					default:
						msg.Text = "Не знаю такой команды"
					}
				}

				if _, resultMenuBookingsStage := userStage.(*stages.BookingsMenuStage); resultMenuBookingsStage {
					switch update.Message.Text {
					case "Забронировать отель":
						newCreateBookingStage := stages.NewCreateBookingStage()
						userInfo.SetUserStage(update.Message.From.ID, newCreateBookingStage)
						newCreateBookingStage.ConfigureMessage(&msg)
					case "Вернутся обратно":
						newMainStage := stages.NewMainStage()
						userInfo.SetUserStage(update.Message.From.ID, newMainStage)
						newMainStage.ConfigureMessage(&msg)
					case "Получить по ID":
						newBookingsGetBookingIdStage := stages.NewBookingsGetBookingIdStage()
						userInfo.SetUserStage(update.Message.From.ID, newBookingsGetBookingIdStage)
						newBookingsGetBookingIdStage.ConfigureMessage(&msg)
					case "Получить по ID отеля":
						newBookingGetHotelIdStage := stages.NewBookingGetHotelIdStage()
						userInfo.SetUserStage(update.Message.From.ID, newBookingGetHotelIdStage)
						newBookingGetHotelIdStage.ConfigureMessage(&msg)
					default:
						msg.Text = "Не знаю такой команды"
					}
				}

				if bookingsGetBookingIdStage, resultBookingsGetBookingIdStage := userStage.(*stages.BookingsGetBookingIdStage); resultBookingsGetBookingIdStage {
					inputError := bookingsGetBookingIdStage.ComputeInput(userInfo, update.Message.From.ID, update.Message.Text)
					if inputError != nil {
						newBookingsMenuStage := stages.NewBookingsMenuStage(*currUser.Role)
						userInfo.SetUserStage(update.Message.From.ID, newBookingsMenuStage)
						newBookingsMenuStage.ConfigureMessage(&msg)
						msg.Text = "Ошибка при вводе данных. " + msg.Text
					} else {
						bookingsGetBookingIdStage.EndStage()
						if bookingsGetBookingIdStage.StagesFinished() {
							newBookingsMenuStage := stages.NewBookingsMenuStage(*currUser.Role)
							userInfo.SetUserStage(update.Message.From.ID, newBookingsMenuStage)
							newBookingsMenuStage.ConfigureMessage(&msg)

							userBooking := userInfo.GetUserBooking(update.Message.From.ID)
							booking, errGetBooking := apiService.GetBookingByID(userBooking.BookingID, currUser)
							if errGetBooking != nil {
								msg.Text = "Ошибка при получении бронирования"
							} else {
								msg.Text = data_representation.GetBooking(booking)
							}
						}
					}
				}

				if bookingsGetHotelIdStage, resultBookingsGetHotelIdStage := userStage.(*stages.BookingsGetHotelIdStage); resultBookingsGetHotelIdStage {
					inputError := bookingsGetHotelIdStage.ComputeInput(userInfo, update.Message.From.ID, update.Message.Text)
					if inputError != nil {
						newBookingsMenuStage := stages.NewBookingsMenuStage(*currUser.Role)
						userInfo.SetUserStage(update.Message.From.ID, newBookingsMenuStage)
						newBookingsMenuStage.ConfigureMessage(&msg)
						msg.Text = "Ошибка при вводе данных. " + msg.Text
					} else {
						bookingsGetHotelIdStage.EndStage()
						if bookingsGetHotelIdStage.StagesFinished() {
							newBookingsMenuStage := stages.NewBookingsMenuStage(*currUser.Role)
							userInfo.SetUserStage(update.Message.From.ID, newBookingsMenuStage)
							newBookingsMenuStage.ConfigureMessage(&msg)

							userBooking := userInfo.GetUserBooking(update.Message.From.ID)
							bookings, errGetBooking := apiService.GetBookings(userBooking.BookingID, currUser)
							if errGetBooking != nil {
								msg.Text = "Ошибка при получении бронирований"
							} else {
								for _, booking := range bookings {
									msg.Text = data_representation.GetBooking(&booking)
									if _, err := bot.Send(msg); err != nil {
										log.Panic(err)
									}
								}
								continue
							}
						}
					}
				}

				if hotelsGetHotelIdStage, resultHotelsGetHotelIdStage := userStage.(*stages.HotelsGetHotelIdStage); resultHotelsGetHotelIdStage {
					inputError := hotelsGetHotelIdStage.ComputeInput(userInfo, update.Message.From.ID, update.Message.Text)
					if inputError != nil {
						newHotelsMenuStage := stages.NewHotelsMenuStage(*currUser.Role)
						userInfo.SetUserStage(update.Message.From.ID, newHotelsMenuStage)
						newHotelsMenuStage.ConfigureMessage(&msg)
						msg.Text = "Ошибка при вводе данных. " + msg.Text
					} else {
						hotelsGetHotelIdStage.EndStage()
						if hotelsGetHotelIdStage.StagesFinished() {
							newHotelsMenuStage := stages.NewHotelsMenuStage(*currUser.Role)
							userInfo.SetUserStage(update.Message.From.ID, newHotelsMenuStage)
							newHotelsMenuStage.ConfigureMessage(&msg)

							userHotel := userInfo.GetUserHotel(update.Message.From.ID)

							hotel, errGetHotel := apiService.GetHotelByID(userHotel.ID, currUser)
							if errGetHotel != nil {
								msg.Text = "Ошибка при получении отеля"
							} else {
								msg.Text = data_representation.GetHotel(hotel)
							}

						}
					}
				}

				if createBookingStage, resultCreateBookingStage := userStage.(*stages.CreateBookingStage); resultCreateBookingStage {
					inputError := createBookingStage.ComputeInput(userInfo, update.Message.From.ID, update.Message.Text)
					if inputError != nil {
						newBookingsMenuStage := stages.NewBookingsMenuStage(*currUser.Role)
						userInfo.SetUserStage(update.Message.From.ID, newBookingsMenuStage)
						newBookingsMenuStage.ConfigureMessage(&msg)
						msg.Text = "Ошибка при вводе данных. " + msg.Text
					} else {
						createBookingStage.EndStage()
						if createBookingStage.StagesFinished() {
							newBookingsMenuStage := stages.NewBookingsMenuStage(*currUser.Role)
							userInfo.SetUserStage(update.Message.From.ID, newBookingsMenuStage)
							newBookingsMenuStage.ConfigureMessage(&msg)

							createResult, errCreate := createBookingStage.Finish(userInfo, update.Message.From.ID, apiService)
							if errCreate != nil {
								msg.Text = "Ошибка при создании бронирования"
							} else {
								if createResult {
									msg.Text = "Успешно создано бронирование"
								} else {
									msg.Text = "Не удалось создать бронирование"
								}
							}
						} else {
							createBookingStage.ConfigureMessage(&msg)
						}
					}
				}

				if _, resultMenuHotelsStage := userStage.(*stages.HotelsMenuStage); resultMenuHotelsStage {
					switch update.Message.Text {
					case "Создать отель":
						newCreateHotelStage := stages.NewCreateHotelStage()
						userInfo.SetUserStage(update.Message.From.ID, newCreateHotelStage)
						newCreateHotelStage.ConfigureMessage(&msg)
					case "Вернутся обратно":
						newMainStage := stages.NewMainStage()
						userInfo.SetUserStage(update.Message.From.ID, newMainStage)
						newMainStage.ConfigureMessage(&msg)
					case "Получить по ID":
						newHotelsGetHotelIdStage := stages.NewHotelsGetHotelIdStage()
						userInfo.SetUserStage(update.Message.From.ID, newHotelsGetHotelIdStage)
						newHotelsGetHotelIdStage.ConfigureMessage(&msg)
					case "Получить все":
						hotels, errHotels := apiService.GetHotels(currUser)
						if errHotels != nil {
							msg.Text = "Ошибка при получении отелей"
						} else {
							for _, hotel := range hotels {
								msg.Text = data_representation.GetHotel(&hotel)
								if _, err := bot.Send(msg); err != nil {
									log.Panic(err)
								}
							}
							continue
						}

					default:
						msg.Text = "Не знаю такой команды"
					}
				}

				if createHotelStage, resultCreateHotelStage := userStage.(*stages.CreateHotelStage); resultCreateHotelStage {
					inputError := createHotelStage.ComputeInput(userInfo, update.Message.From.ID, update.Message.Text)
					if inputError != nil {
						newHotelsMenuStage := stages.NewHotelsMenuStage(*currUser.Role)
						userInfo.SetUserStage(update.Message.From.ID, newHotelsMenuStage)
						newHotelsMenuStage.ConfigureMessage(&msg)
						msg.Text = "Ошибка при вводе данных. " + msg.Text
					} else {
						createHotelStage.EndStage()
						if createHotelStage.StagesFinished() {
							newHotelsMenuStage := stages.NewHotelsMenuStage(*currUser.Role)
							userInfo.SetUserStage(update.Message.From.ID, newHotelsMenuStage)
							newHotelsMenuStage.ConfigureMessage(&msg)
							createResult, errCreate := createHotelStage.Finish(userInfo, update.Message.From.ID, apiService)
							if errCreate != nil {
								msg.Text = "Ошибка при создании отеля"
							} else {
								if createResult {
									msg.Text = "Успешно создан отель"
								} else {
									msg.Text = "Не удалось создать отель"
								}
							}
						} else {
							createHotelStage.ConfigureMessage(&msg)
						}
					}
				}

			} else {
				continue
			}
		}
		// sending message
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
