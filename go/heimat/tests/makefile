
dev: 
	@modd --file=modd-dev.conf

doc:
	godoc -http=:6060 -play=true


tier1:

	@clear
	@go test \
	-v \
	-run "Tier1"\
	./src/... \



tier2:
	$(eval export maxtier := 2)
	make _tier

tier3:
	$(eval export maxtier := 3)
	make _tier

tier4:
	$(eval export maxtier := 4)
	make _tier

tier5:
	$(eval export maxtier := 5)
	make _tier

tier6:
	$(eval export maxtier := 6)
	make _tier

tier7:
	$(eval export maxtier := 7)
	make _tier

tier8:
	$(eval export maxtier := 8)
	make _tier

tier9:
	$(eval export maxtier := 9)
	make _tier


_tier:
	@clear
	@go test \
	-v \
	-count=1 \
	-run "Tier[1-${maxtier}]" \
	./src/... \
