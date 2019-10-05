package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"bufio"
	"github.com/spf13/pflag"
)

type selpg_args struct{
	stpg int
	edpg int
	input_filename string
	pglen int	
	pgtype string	
	print_dest string	
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, os.Args[0]+" -sstart_page_num -eend_page_num [ -f | -llines_per_page ] [ -ddest ] [ input_filename ]\n")
    pflag.PrintDefaults()
}

func Init(args *selpg_args){
	pflag.Usage = usage
    pflag.IntVarP(&(args.stpg), "start", "s", -1, "start page")
    pflag.IntVarP(&(args.edpg), "end", "e", -1, "end page")
	pflag.IntVarP(&(args.pglen), "line", "l", 72, "page lines")
	pflag.StringVarP(&(args.pgtype), "type", "f", "l", "type of print")
    pflag.StringVarP(&(args.print_dest), "destination", "d", "", "print destination")
}

//Check whether parameters are valid or not
func check(args *selpg_args){
	//Not enough args
	if len(os.Args) < 3 || args.stpg == -1 || args.edpg == -1{
		fmt.Fprintf(os.Stderr, "ERROR: %s not enough parameters!\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Input the parameters start page and end page!\n")
		usage()
		os.Exit(0)
	}

	//Page's index error
	if args.stpg <= 0{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid start page number! Can not less than or equal to 0!\n")
		os.Exit(1)
	}
	if args.edpg <= 0{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid end page number! Can not less than or equal to 0!\n")
		os.Exit(1)
	}
	if args.stpg > math.MaxInt32 - 1{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid start page number! Can not grater than %v!\n",math.MaxInt32-1)
		os.Exit(1)
	}
	if args.edpg > math.MaxInt32 - 1{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid end page number! Can not grater than %v!\n",math.MaxInt32-1)
		os.Exit(1)
	}
	if args.edpg < args.stpg{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid end page number! Can not less than start page number!\n")
		os.Exit(1)
	}

	//pglen error
	if args.pglen <= 0{
		fmt.Fprintf(os.Stderr, "ERROR: The pages' length can not be less than or equal to 0!\n")
		os.Exit(1)
	}
	if args.pglen > math.MaxInt32 - 1{
		fmt.Fprintf(os.Stderr, "ERROR: The pages' length can not be grater than %v!\n",math.MaxInt32-1)
		os.Exit(1)
	}

	//page type error
	if args.pgtype != "f" && args.pgtype != "l"{
		fmt.Fprintf(os.Stderr,"%sthis\n",args.pgtype);
		fmt.Fprintf(os.Stderr, "ERROR: thisThe page type must be 'f' or 'l'!\n")
		os.Exit(1)
	}
}

func process_input(args *selpg_args){
	var fin *os.File
	if args.input_filename != ""{
		var inErr error
		fin, inErr = os.Open(args.input_filename)
		if inErr != nil{
			fmt.Fprintf(os.Stderr, "\nERROR! Can not open the input file: %s\n", args.input_filename)
			os.Exit(1)
		}
	}else{
		fin = os.Stdin
	}
	finBuffer := bufio.NewReader(fin)


	var fout io.WriteCloser
	var cmd *exec.Cmd
	if args.print_dest != ""{
		cmd = exec.Command("lp", "-d", args.print_dest)
		var outErr error
		cmd.Stdout, outErr = os.OpenFile(args.print_dest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		fout, outErr = cmd.StdinPipe()
		if outErr != nil{
			fmt.Fprintf(os.Stderr, "\nERROR! Can not open pipe to \"lp -d%s\"\n",  args.print_dest)
			os.Exit(1)
		}
		cmd.Start()
		cmd.Wait()
	}else{
		fout = os.Stdout
	}
	
	/* begin one of two main loops based on page type */
	if args.pgtype == "l" {
		line_ctr := 0
		page_ctr := 1
		for {
			line,  crc := finBuffer.ReadString('\n')
			if crc != nil {
				break
			}
			line_ctr ++
			if line_ctr > args.pglen {
				page_ctr ++
				line_ctr = 1
			}
	
			if (page_ctr >= args.stpg) && (page_ctr <= args.edpg) {
				_, err := fout.Write([]byte(line))
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERROR!", err)
					os.Exit(1)
				}
			 }
		}  
	}else{
		page_ctr := 1
		for{
			
			page,  crc := finBuffer.ReadString('\f')
			if crc != nil {
				break
			}
			page_ctr ++
			if ( (page_ctr >= args.stpg) && (page_ctr <= args.edpg) ){
				_, err := fout.Write([]byte(page))
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERROR!", err)
					os.Exit(1)
				}
			}
		}
	}

	fin.Close()
	fout.Close()
}

func main() {
	args := new(selpg_args)
	Init(args)
    pflag.Parse()
    
    othersArg := pflag.Args()
    if pflag.NArg() > 0 {
        args.input_filename = othersArg[0]
    } else {
        args.input_filename = ""
	}

	check(args)
	process_input(args)
}