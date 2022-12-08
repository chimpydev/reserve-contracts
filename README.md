# Protocol overview
Reserve Insurance is a protocol to insure(hedge) against pegged assets and speculate on depegs.  
The aim of this protocol is to allow users to HEDGE & RISK these insurances, also providing yield to incentivize early depositors on both side of insurance (buyers and sellers).  
We consider insurance sellers to be the same as Risk users, or collateral depositors. We also consider insurance buyers to be the same as Hedge users, or premium depositors.  
This is made possible by having Vaults that leverage ERC4626 and a Chainlink Oracle to determine the current price.  
When the price of a pegged asset is below the strike price of a Vault, a Keeper(could be anyone) will trigger the depeg event and both Vaults(hedge and risk) will swap their total assets with the other party.  
These insurances last a certain amount of time, and if a depeg event does not happen between that time period, a Keeper will trigger the end of epoch event, being that the Hedge Vault then has to give all of its total assets to the Risk Vault.
It is issued a token representing the depositors position which is 1 to 1 ratio'ed to the deposited amount, and users can deposit that token in a fork of Synthetix Rewards to "yield farm" governance tokens. 

# Smart Contracts
 
### Controller.sol -   
   
Smart Contract used to call Oracles' latest price data, and trigger epoch end or depeg events. They can be triggered if the Id of that Market/Vault epoch is less then the current block.timestamp, or if the current price of that insured pegged asset is below the Vaults' Strike Price.
These functions will then call that Market, figure out the pair of Vaults of that Market, and transfer the underlying assets between the Risk/Hedge parties. Must mention that inside this contract we convert all price feed results to 18 decimal points, so we can compare to the Strike Price passed by the admin on "VaultFactory.sol"
> Responsibility:  
- Trigger Depegs;  
- Trigger end of Epoch;  
- Check Oracle Price Feed converted to 18 decimals;  

### VaultFactory.sol -   
  
Factory for the creation of Vault pairs. A pair of Vaults constructs a Market, a Hedge and a Risk Vault. As said before, Markets are determined by its Token Insured & Strike Price, and have an Index associated with it.  Since Markets have Vaults that always come in pairs, some functions use that assumption for their logic, for example arrays that return the vault addresses have in the [0] position the Hedge Vault and on the [1] position the Risk Vault.
This contract allows the Admin role to change params inside the Vault, and also the Oracle associated with the insured token.  
Must mention that we expect admin to pass the StrikePrice as a 18 decimal int, so we can compare to the Oracle Price using the same decimals.
> Responsability:  
- Create pair of Vaults (Markets);
- Change Variables;  

### rewards/RewardsFactory.sol -   
   
Contract used to deploy the pair of "StakingRewards.sol" to reflect the Hedge/Risk vaults implementation above.  
Must mention we used a keccak256 to hash a market index from "VaultFactory.sol" and its "Vault.sol" respective market epochEnd/id to represent the StakingRewards vaults as an index to be able to query it later.
> Responsibility:  
- Create pair of Farms from Markets;


# Concerns & Additional Information  
We consider the name of the Token for frontend uses, to be in this format "depegUSDC_99*" , this has both insured token name in CAPS and strike Price. 