package api

import(
        "context"
        "net"
        "log"

	"headline/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/grpc/credentials/insecure"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
        db *gorm.DB
)

func InitDb(dsi string) error {
        var err error
        db, err = gorm.Open(postgres.Open(dsi))

        return err
}

func AutoMigrate() error {
        return db.AutoMigrate(&model.Article{}, &model.Interest{})
}

type Connection struct {
        server          *grpc.Server
        clientConn      *grpc.ClientConn
        closer          func()
        startup         func()
}

func GenerateTestServer(ctx context.Context) (*Connection, error) {
        buffer := 101024 * 1024
        lis := bufconn.Listen(buffer)
        
        s := grpc.NewServer()

        start := func() {
                if err := s.Serve(lis); err != nil {
                        log.Printf("Error serving server: %v", err)
                }
        }

        conn, err := grpc.DialContext(ctx, "", 
                grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
                        return lis.Dial()
                }), grpc.WithTransportCredentials(insecure.NewCredentials()))

        if err != nil {
                log.Printf("Error connecting to server: %v", err)
        }

        closer := func() {
                err := lis.Close()

                if err != nil {
                        log.Printf("Error closing listener: %v", err)
                }

                s.Stop()
        }

        return &Connection{
                clientConn: conn,
                closer: closer,
                server: s,
                startup: start,
        }, nil
}

