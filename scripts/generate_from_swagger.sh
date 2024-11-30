#!/bin/bash


swagger generate server -f booking/api/swagger/booking.yaml -t booking/internal --exclude-main --principal models.User

swagger generate server -f hotel/api/swagger/hotel.yaml -t hotel/internal --exclude-main --principal models.User

swagger generate server -f auth/api/swagger/auth.yaml -t auth/internal --exclude-main

echo "REGENERATED. NOW TIDYING"

cd booking || exit
go mod tidy

cd ../hotel || exit
go mod tidy

cd ../auth || exit
go mod tidy

