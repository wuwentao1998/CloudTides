import React from "react";
import PropTypes from "prop-types";
import {Container, Row, Col} from "shards-react";

import PageTitle from "./../components/common/PageTitle";
import PoliciesTable from "../components/contribution/PoliciesTable";

const Contribution = () => (
    <Container fluid className="main-content-container px-4">
        {/* Page Header */}
        <Row noGutters className="page-header py-4">
            <PageTitle title="Contribution" subtitle="Manage Policies" className="text-sm-left mb-3"/>
        </Row>
        <Row>
            {/* Resources */}
            <Col className="col-lg mb-4">
                <PoliciesTable/>
            </Col>
        </Row>
    </Container>
);

Contribution.propTypes = {};

Contribution.defaultProps = {};

export default Contribution;
