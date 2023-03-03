import React from "react"
import './navbar.component.css'
import {Container, Dropdown, Image, Nav, Navbar, NavDropdown} from "react-bootstrap";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {faUser, faUserPen, faSignOut, faUserCircle} from '@fortawesome/free-solid-svg-icons'
import UserProfile from "./modal/userprofile";
import ChangePassword from "./modal/changepassword";
import Swal from "sweetalert2";

export const NavigationBar = () => {
    const [modalUserProfileShow, setModalUserProfileShow] = React.useState(false);
    const [modalChangePasswordShow, setModalChangePasswordShow] = React.useState(false);
    const handleSignOut = () => {
        Swal.fire({
            title: 'You are about to sign out',
            text: "Are you sure?",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes'
        }).then((result) => {
            if (result.isConfirmed) {
                localStorage.removeItem("user")
                localStorage.removeItem("mode")
                localStorage.removeItem("nik")
                localStorage.removeItem("sidebar")
                Swal.fire({
                    icon: 'success',
                    title: 'You already signed out',
                    showConfirmButton: false,
                    timer: 1500
                }).then(() => window.location.href = "/")
            }
        })
    }
    return (
        <Navbar collapseOnSelect expand="lg" className="px-5">
            <Container fluid>
                <h1 className="text-white">SIPK</h1>
                <Dropdown className="ms-auto">
                    <Dropdown.Toggle id="dropdown-basic" className="d-flex align-items-center flex-row gap-2 bg-transparent border-0">
                        <FontAwesomeIcon icon={faUserCircle} style={{fontSize: "2rem"}}/>
                        <span>Julyus Andreas</span>
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                        <Dropdown.Item href="" onClick={() => setModalUserProfileShow(true)}>
                            <FontAwesomeIcon icon={faUser} className="pe-2"/> View Profile
                        </Dropdown.Item>
                        <Dropdown.Item href="" onClick={() => setModalChangePasswordShow(true)}>
                            <FontAwesomeIcon icon={faUserPen} className="pe-2"/> Change Password
                        </Dropdown.Item>
                        <Dropdown.Divider />
                        <Dropdown.Item onClick={() => handleSignOut()}>
                            <FontAwesomeIcon icon={faSignOut} className="pe-2"/> Log out
                        </Dropdown.Item>
                    </Dropdown.Menu>
                </Dropdown>
            </Container>
            <UserProfile
                show={modalUserProfileShow}
                onHide={() => setModalUserProfileShow(false)}
            />
            <ChangePassword
                show={modalChangePasswordShow}
                onHide={() => setModalChangePasswordShow(false)}
            />
        </Navbar>
    )
}
