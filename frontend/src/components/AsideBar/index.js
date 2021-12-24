import { AsideBarContainer } from "./styled";
import { Link } from "react-router-dom";
import { ProductContext } from "../../context/ProductContext";
import { useContext } from "react";

export default function AsideBar() {
  const { products } = useContext(ProductContext);

  return (
    <AsideBarContainer>
      <div>
        <h3>Client Name</h3>
        <p>{products.length} produtos encontrados</p>
      </div>

      <nav>
        <Link to="/">Home</Link>
        <Link to="/cart">Carrinho</Link>
        <Link to="/credits">Creditos</Link>
      </nav>
    </AsideBarContainer>
  );
}
