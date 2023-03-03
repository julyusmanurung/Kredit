import './footer.css'
import {Container} from "react-bootstrap";

export const Footer = () => {
    return(
        <Container fluid className="footer px-5 py-3">
            <label className="text-white">Copyright © 2023 <strong>Bank Sinarmas</strong>. All rights reserved.</label>
        </Container>
    )
}

