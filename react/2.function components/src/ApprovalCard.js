import React from 'react';

const ApprovalCard = (props) => {
    return(
        <div className="ui card">
          <div className="content">{props.children} </div>
          <div className="extra content">
                <div class="ui two buttons">
                    <div class="ui basic green button">Approve</div>
                    <div class="ui basic red button">Decline</div>
                </div>
            </div>
        </div>
    );
};
//props.children就是抓取包在<ApprovalCard></ApprovalCard>裡面的數值

export default ApprovalCard;