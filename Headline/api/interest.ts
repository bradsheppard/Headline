import {InterestServiceClient} from "../proto/InterestServiceClientPb";
import {GetInterestsRequest} from "../proto/interest_pb";

class InterestService {
    static async getInterests(userId: number): Promise<Array<string>> {
        const interestServiceClient = new InterestServiceClient('http://envoy:9090')

        const request = new GetInterestsRequest()
        request.setUserid(userId)

        const response = await interestServiceClient.getInterests(request, null)
        return response.getInterestsList().map(x => x.getName())
    }
}

export default InterestService;

