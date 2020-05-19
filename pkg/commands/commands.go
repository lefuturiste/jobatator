package commands

// CmdMap -
var CmdMap = []CmdDefinition{
	{
		Name:        "PING",
		Handler:     Ping,
		RequireAuth: false,
		UseGroup:    false,
		Description: "Check if the connexion is alive",
	}, {
		Name:        "QUIT",
		Handler:     Quit,
		RequireAuth: false,
		UseGroup:    false,
		Description: "Disconnect from the server",
	}, {
		Name:        "AUTH",
		Handler:     Auth,
		RequireAuth: false,
		UseGroup:    false,
		Args:        2,
		Usage:       "AUTH {username} {password}",
		Description: "Authenticate with a username and a password",
	}, {
		Name:        "USE_GROUP",
		Handler:     UseGroup,
		RequireAuth: true,
		UseGroup:    false,
		Args:        1,
		Usage:       "USE_GROUP {group}",
		Description: "Declare which group the client want to use",
	}, {
		Name:        "DEBUG",
		Handler:     Debug,
		RequireAuth: false,
		UseGroup:    false,
		Description: "Dump server data in a human readable format, only work if debug is enabled",
	}, {
		Name:        "DEBUG_JSON",
		Handler:     DebugJSON,
		RequireAuth: false,
		UseGroup:    false,
		Description: "Dump server data in a JSON format, only work if debug is enabled",
	}, {
		Name:        "STOP_SERVER",
		Handler:     StopServer,
		RequireAuth: false,
		UseGroup:    false,
		Description: "Stop the server, only work if test_mode is enabled",
	}, {
		Name:        "PUBLISH",
		Handler:     Publish,
		RequireAuth: true,
		UseGroup:    true,
		Args:        3,
		Usage:       "PUBLISH {queue_name} {job_type} {job_payload}",
		Description: "Create a job in a queue and return the ID of the created job",
	}, {
		Name:        "SUBSCRIBE",
		Handler:     Subscribe,
		RequireAuth: true,
		UseGroup:    true,
		Args:        1,
		Usage:       "SUBSCRIBE {queue_name}",
		Description: "Subscribe to a queue to receive jobs",
	}, {
		Name:        "LIST_QUEUES",
		Handler:     ListQueues,
		RequireAuth: true,
		UseGroup:    true,
		Description: "List all queues in the current group",
	}, {
		Name:        "DELETE_QUEUE",
		Handler:     DeleteQueue,
		RequireAuth: true,
		UseGroup:    true,
		Args:        1,
		Usage:       "DELETE_QUEUE {queue_name}",
		Description: "Delete a queue",
	}, {
		Name:        "LIST_JOBS",
		Handler:     ListJobs,
		RequireAuth: true,
		UseGroup:    true,
		Args:        1,
		Usage:       "LIST_JOBS {queue_name}",
		Description: "List all jobs in a queue",
	}, {
		Name:        "UPDATE_JOB",
		Handler:     UpdateJob,
		RequireAuth: true,
		UseGroup:    true,
		Args:        2,
		Usage:       "UPDATE_JOB {job_id} {job_status}",
		Description: "Set the status of a job",
	}, {
		Name:        "DELETE_JOB",
		Handler:     DeleteJob,
		RequireAuth: true,
		UseGroup:    true,
		Args:        1,
		Usage:       "DELETE_JOB {job_id}",
		Description: "Delete a job",
	}, {
		Name:        "DISPATCH",
		Handler:     Dispatch,
		RequireAuth: true,
		UseGroup:    true,
		Description: "Will force the dispatch of jobs accross all the workers",
	},
}