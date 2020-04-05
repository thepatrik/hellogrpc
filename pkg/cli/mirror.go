package cli

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/thepatrik/hellogrpc/internal/pb/mirror"

	"github.com/spf13/cobra"
	"github.com/thepatrik/strcolor"
	"google.golang.org/grpc"
)

var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Mirror text",
	Long:  "Mirror text",
	Run:   mirror,
}

func mirror(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetString("host")
	if host == "" {
		showHelp(cmd, "missing host")
	}

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := pb.NewMirrorClient(conn)

	text := "The quick brown 狐 jumped over the lazy 犬"

	fmt.Printf("%-20s%s\n", strcolor.Magenta("input:"), text)

	output, err := mirrorText(text, client)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%-20s%s\n", strcolor.Magenta("output:"), strcolor.Cyan(output))

	for {
		fmt.Printf("%-20s", strcolor.Magenta("input:"))
		text = read()
		if text == "" {
			break
		}

		output, err := mirrorText(text, client)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("%-20s%s\n", strcolor.Magenta("output:"), strcolor.Cyan(output))
	}
}

func mirrorText(text string, client pb.MirrorClient) (string, error) {
	req := &pb.MirrorTextRequest{Text: text}
	res, err := client.MirrorText(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.Text, nil
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

func init() {
	mirrorCmd.Flags().String("host", "localhost:9090", "Host address")
	AddCommand(mirrorCmd)
}
