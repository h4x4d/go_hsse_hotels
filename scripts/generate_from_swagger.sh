#!/bin/bash


swagger generate server -f booking/api/swagger/booking.yaml -t booking/internal
rm -rf booking/internal/cmd

swagger generate server -f hotel/api/swagger/hotel.yaml -t hotel/internal
rm -rf hotel/internal/cmd

echo "REGENERATED. NOW TIDYING"

cd booking || exit
go mod tidy

cd ../hotel || exit
go mod tidy
