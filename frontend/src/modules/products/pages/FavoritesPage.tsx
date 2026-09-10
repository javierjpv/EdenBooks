import { Container } from "@mui/material"
import { FavoriteList } from "../components/FavoriteList"

export const FavoritesPage = () => {
  return (
    <>
    <Container sx={{ marginTop: 22 }}>
      <h1>My ProductPage</h1>
      <FavoriteList  />
    </Container>
  </>
  )
}
