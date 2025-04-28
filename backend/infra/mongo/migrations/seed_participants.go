package migrations

import (
	"context"

	"github.com/azevedoMairon/decidr-app/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedParticipants(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection("participants")

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	participants := []interface{}{
		entities.Participant{Name: "Aline", ImageUrl: "https://s2.glbimg.com/BP82Ew-JnAgxPxFnvIhvdqaLsPs=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/6/K/3rOk79SQiWAqL0AnZxJw/aline-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Vinícius", ImageUrl: "https://s2.glbimg.com/HY3hReI2ddPdQbzfutbdIIoe8eQ=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/P/7/OTwcPASPGP4gbmy3QsuA/vinicius-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Arleane", ImageUrl: "https://s2.glbimg.com/cig7f5qU-1V2WE14Rj-MHsqxwMo=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/A/u/ncRjVERJAv9VjvMpiQyA/arleane-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Marcelo", ImageUrl: "https://s2.glbimg.com/3reaY-P_hhUQkJnj3hzaj4H2V_g=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/0/T/sBh6wlRx6CBU2QX7m6lA/marcelo-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Camila", ImageUrl: "https://s2.glbimg.com/2x_1MSjYLO1zaHfHjXnM3RlRdtM=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/w/R/5qAqGiTEArB0yMUptk5A/camila-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Thamiris", ImageUrl: "https://s2.glbimg.com/PY0BHk7p8bhg2VIimyKS24d1P04=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/i/S/At5qCZSz6EqHtvBTJHLw/thamiris-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Daniele Hypolito", ImageUrl: "https://s2.glbimg.com/tMhq3B0Tz4cukRi0AMKyr5eeoGw=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/h/2/22pebxTkGxchRBUlK8MQ/daniele-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Diego Hypolito", ImageUrl: "https://s2.glbimg.com/zMwedNYl83ucrGEep3QspGvC2NM=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/L/M/TWeMYUTx2BTlJM5qRKtQ/diego-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Diogo Almeida", ImageUrl: "https://s2.glbimg.com/G96s9uUh5NelnbCSCfXLXE6qoeQ=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/b/y/RwiHQzSkWIbm5YIHga0w/diogo-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Vilma", ImageUrl: "https://s2.glbimg.com/b5Q-1gz6YZZR2osc7Z6IHIgg0Qo=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/Q/B/prKIZiQAGkJPRtHrRirw/vilma-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Edilberto", ImageUrl: "https://s2.glbimg.com/T2-nzU8RGQYWt1jWUGofUYwWkoE=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/l/5/P8dJ0BSPmTeCeeb0Caiw/edilberto-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Raissa", ImageUrl: "https://s2.glbimg.com/07z_KxTItCLUX9a9QerzK4RJYXE=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/D/t/raObCaTgA5niMPugBMKg/raissa-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Eva", ImageUrl: "https://s2.glbimg.com/iWVO7h3QlGiV8BwD0RYDx7cEwxU=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/L/B/wM4kF4Tpy0SpX2sA4rwQ/eva-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Renata", ImageUrl: "https://s2.glbimg.com/NWPDpEC0-wPrndUxzfLjoCqRUcw=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/t/r/zVRCCJQCivrP42Wnlstg/renata-bbb-25.png", IsNominated: true, IsEliminated: false},
		entities.Participant{Name: "Gabriel", ImageUrl: "https://s2.glbimg.com/0xNq2yiP56tVmQTrsJugX9F16Ds=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/G/u/PtDQNWQvat18KX5MrbxQ/gabriel-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Maike", ImageUrl: "https://s2.glbimg.com/kT69XDgS68cKcD2oQx11OBpwZFg=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/3/e/DcDKDsT1KFSELrDBP8CQ/maike-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Gracyanne Barbosa", ImageUrl: "https://s2.glbimg.com/-dBPW_G7xXmtTVTY4He-CsGqiR4=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/O/b/2aAkGqQVye0oLJBQNpjw/gracyanne-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Giovanna", ImageUrl: "https://s2.glbimg.com/cCqbmUTyM5JkF81LA0FDA7UOc7U=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/W/8/lDhnWkT1G14vIqdGDA1Q/giovanna-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "Guilherme", ImageUrl: "https://s2.glbimg.com/DmsApwFL9BbGa8mJVI38sFGOMVY=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/z/F/MAL0oiQWy522aXB9Q0tw/guilherme-bbb-25.png", IsNominated: false, IsEliminated: false},
		entities.Participant{Name: "Joselma", ImageUrl: "https://s2.glbimg.com/kOXnDmx4ZRe9BKv-PI_tYdZutGM=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/a/l/UwBn2JQj2glsBZN4TyiA/joselma-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "João Gabriel", ImageUrl: "https://s2.glbimg.com/L6FYUD9bMsw-P--MJZbDuhjudQQ=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/1/7/LlAxWeTiOwAmMwPnJAZw/joaogabriel-bbb-25.png", IsNominated: false, IsEliminated: true},
		entities.Participant{Name: "João Pedro", ImageUrl: "https://s2.glbimg.com/jXyWF4ADoBKr5iDoc1cVX1IC_dw=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/4/V/exCOZxTdy475FMuDRk3Q/joaopedro-bbb-25.png", IsNominated: true, IsEliminated: false},
		entities.Participant{Name: "Vitória Strada", ImageUrl: "https://s2.glbimg.com/PoantXfakW-hCaMT8vwfzWniQDM=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/6/r/RL6bRLR2yrvZTOHA0E7Q/vitoria-bbb-25.png", IsNominated: true, IsEliminated: false},
		entities.Participant{Name: "Mateus", ImageUrl: "https://s2.glbimg.com/4dDnO_vHi9lRvQgckLqIB-cIaOc=/i.s3.glbimg.com/v1/AUTH_e84042ef78cb4708aeebdf1c68c6cbd6/internal_photos/bs/2025/A/F/5evWSESmuY2k9WwG9rFg/mateus-bbb-25.png", IsNominated: false, IsEliminated: true},
	}

	_, err = collection.InsertMany(ctx, participants)
	return err
}
