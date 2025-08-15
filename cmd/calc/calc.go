package calc


var CalculateCmd = &cobra.Command{
	Use:  "calc [search] ",
	Args: cobra.MinimumNArgs(1),

