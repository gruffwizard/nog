package cli

import ()

type AConfig struct {
	args []string
	envs []string
	foo. []string
}

func (c *AConfig) AddEnv(k string,v string) {
				c.envs = append(c.envs, k+"="+v)
}

func (l *AConfig) Display() {

	/*
	   if l.QuickStart!="" {   fmt.Printf(" QuickStart  : %v\n",l.QuickStart)}
	   if config.SrcDir!=""     {   fmt.Printf(" Source Dir  : %v\n",config.SrcDir)}
	   if config.SrcVol!=""     {   fmt.Printf(" Source Vol  : %v\n",config.SrcVol)}
	   if config.MvnDir!=""     {   fmt.Printf(" Maven Dir   : %v\n",config.MvnDir)}
	   if config.MvnVol!=""     {   fmt.Printf(" Maven Vol   : %v\n",config.MvnVol)}
	   if l.image!=""      {   fmt.Printf(" Image       : %v\n",l.image)}
	   if len(l.Args)>0    {   fmt.Printf(" Args        : %v\n",l.Args)}
	   if len(l.cmd)>0     {   fmt.Printf(" Command     : %v\n",l.cmd)}

	*/

}
