import React from "react";
import PropTypes from "prop-types";
import {Container, Row, Col} from "shards-react";

import MainNavbar from "../components/layout/MainNavbar/MainNavbar";
import MainSidebar from "../components/layout/MainSidebar/MainSidebar";
import MainFooter from "../components/layout/MainFooter";


class DefaultLayout extends React.Component {
    render() {
        const {noNavbar, noFooter, children} = this.props;
        const offset = (this.props.location.pathname !== "/manage-resources") ? {marginLeft: "80px"} : {};
        return (
            <Container fluid>
                <Row>
                    <MainSidebar {...this.props}/>
                    <Col
                        className="main-content p-0 bg-white"
                        // lg={{size: 10, offset: 2}}
                        // md={{size: 9, offset: 3}}
                        // sm="12"
                        tag="main"
                        style={offset}
                    >
                        {!noNavbar && <MainNavbar/>}
                        {children}
                        {!noFooter && <MainFooter/>}
                    </Col>
                </Row>
            </Container>
        )
    }
}

DefaultLayout.propTypes = {
    /**
     * Whether to display the navbar, or not.
     */
    noNavbar: PropTypes.bool,
    /**
     * Whether to display the footer, or not.
     */
    noFooter: PropTypes.bool
};

DefaultLayout.defaultProps = {
    noNavbar: false,
    noFooter: false
};

export default DefaultLayout;
