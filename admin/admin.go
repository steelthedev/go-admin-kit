package admin

type Admin struct {
  database Database
  router Router
}

func NewAdmin (db Database, router Router) *Admin {
  return &Admin {
    db,
    router
  }
}