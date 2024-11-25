#!/bin/bash


swagger generate server -f booking/api/swagger/booking.yaml -t booking/internal
rm -rf booking/internal/cmd

swagger generate server -f hotel/api/swagger/hotel.yaml -t hotel/internal
rm -rf hotel/internal/cmd

swagger generate server -f auth/api/swagger/auth.yaml -t auth/internal
rm -rf auth/internal/cmd

echo "REGENERATED. NOW TIDYING"

cd booking || exit
go mod tidy

cd ../hotel || exit
go mod tidy

cd ../auth || exit
go mod tidy

