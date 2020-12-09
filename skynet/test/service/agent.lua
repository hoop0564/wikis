skynet.start(function ()
	print("recv a connection", clientid, addr)
	rds[1] = skynet.uniqueservice("redis")
	hall = skynet.uniqueservice("hall")
	socket.start(clientid) -- 绑定 fd actorid 与 epoll事件
	skynet.fork(process_socket_events)
	skynet.dispatch("lua", function (_, _, cmd, ... )
		if cmd == "game_over" then
			game_over()
	end)
end)