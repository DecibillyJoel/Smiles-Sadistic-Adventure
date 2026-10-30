return {
	_lookUp = {},
	startTime = socket.gettime(),
	previousTime = socket.gettime(),
	currentTime = socket.gettime(),
	deltaTime = 0,
	init = function(self)
		self._lookUp[go.get_id()] = true
		return self.currentTime
	end,
	update = function(self)
		if self._lookUp[go.get_id()] then
			self._lookUp = {}
			self.previousTime = self.currentTime
			self.currentTime = socket.gettime()
			self.deltaTime = (self.currentTime ~= self.previousTime) and (self.currentTime - self.previousTime) or false
		end
		self._lookUp[go.get_id()] = true
		return self.deltaTime
	end
}